package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dsoprea/go-exif/v3"
	common "github.com/dsoprea/go-exif/v3/common"
)

var (
	extImagen = map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".tif": true,
		".tiff": true, ".webp": true, ".bmp": true,
	}
	extVideo = map[string]bool{
		".mp4": true, ".mov": true, ".avi": true, ".mkv": true,
		".webm": true, ".m4v": true,
	}
)

// ID de etiquetas EXIF (IFD0) que nos interesan.
const (
	tagMake             = 0x010F
	tagModel            = 0x0110
	tagDateTime         = 0x0132
	tagDateTimeOriginal = 0x9003
	tagArtist           = 0x013B // nombre de usuario / autor si existe
)

func TipoDeArchivo(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	if extImagen[ext] {
		return "imagen"
	}
	if extVideo[ext] {
		return "video"
	}
	return "desconocido"
}

// ExtraerExif lee los metadatos EXIF de un archivo y devuelve un FileMeta.
// Un archivo sin EXIF no es error: se devuelve con los campos vacíos.
func ExtraerExif(path string) (*FileMeta, error) {
	fm := &FileMeta{
		Archivo: filepath.Base(path),
		Tipo:    TipoDeArchivo(path),
	}

	datos, err := os.ReadFile(path)
	if err != nil {
		return fm, nil
	}
	tiff := localizarBloqueTiff(datos)
	if tiff == nil {
		return fm, nil
	}

	ifdMapping, err := common.NewIfdMappingWithStandard()
	if err != nil {
		return fm, err
	}
	tagIndex := exif.NewTagIndex()
	_, index, err := exif.Collect(ifdMapping, tagIndex, tiff)
	if err != nil {
		return fm, err
	}
	root := index.RootIfd
	if root == nil {
		return fm, nil
	}

	// Dispositivo: marca y modelo.
	if marca := valorTag(root, tagMake); marca != "" {
		if modelo := valorTag(root, tagModel); modelo != "" {
			fm.Dispositivo = &Device{Marca: marca, Modelo: modelo}
		} else {
			fm.Dispositivo = &Device{Marca: marca}
		}
	} else if modelo := valorTag(root, tagModel); modelo != "" {
		fm.Dispositivo = &Device{Modelo: modelo}
	}

	// Fecha y hora: DateTimeOriginal primero, DateTime como respaldo.
	if fecha := valorTag(root, tagDateTimeOriginal); fecha != "" {
		fm.FechaHora = fecha
	} else if fecha := valorTag(root, tagDateTime); fecha != "" {
		fm.FechaHora = fecha
	}

	// Usuario: tag Artist si existe.
	if usuario := valorTag(root, tagArtist); usuario != "" {
		fm.Usuario = usuario
	}

	// Ubicacion: GPS (se lee en el sub-IFD GPS, hijo de IFD0).
	var gi *exif.GpsInfo
	for _, child := range root.Children() {
		if g, err := child.GpsInfo(); err == nil && g != nil {
			gi = g
			break
		}
	}
	if gi != nil {
		loc := &Location{}
		tiene := false
		if gi.Latitude.Degrees != 0 || gi.Latitude.Minutes != 0 || gi.Latitude.Seconds != 0 {
			loc.Lat = gi.Latitude.Decimal()
			tiene = true
		}
		if gi.Longitude.Degrees != 0 || gi.Longitude.Minutes != 0 || gi.Longitude.Seconds != 0 {
			loc.Lon = gi.Longitude.Decimal()
			tiene = true
		}
		if gi.Altitude != 0 {
			a := float64(gi.Altitude)
			loc.Alt = &a
			tiene = true
		}
		if tiene {
			fm.Ubicacion = loc
		}
	}

	return fm, nil
}

// localizarBloqueTiff busca el inicio de un bloque TIFF/EXIF dentro de los
// bytes de un archivo. Reconoce la firma "Exif\0\0" (envuelta en JPEG) o la
// firma TIFF directa "II*" / "MM*". Devuelve nil si no hay bloque.
func localizarBloqueTiff(datos []byte) []byte {
	exifSig := []byte("Exif\x00\x00")
	if i := bytes.Index(datos, exifSig); i >= 0 {
		inicio := i + len(exifSig)
		if esCabeceraTiff(datos[inicio:]) {
			return datos[inicio:]
		}
	}
	for _, sig := range [][]byte{{0x49, 0x49, 0x2A, 0x00}, {0x4D, 0x4D, 0x00, 0x2A}} {
		if i := bytes.Index(datos, sig); i >= 0 {
			return datos[i:]
		}
	}
	return nil
}

func esCabeceraTiff(b []byte) bool {
	if len(b) < 4 {
		return false
	}
	return (b[0] == 0x49 && b[1] == 0x49 && b[2] == 0x2A && b[3] == 0x00) ||
		(b[0] == 0x4D && b[1] == 0x4D && b[2] == 0x00 && b[3] == 0x2A)
}

// valorTag devuelve el valor formateado de la primera entrada con ese ID.
func valorTag(root *exif.Ifd, id uint16) string {
	entradas, err := root.FindTagWithId(id)
	if err != nil || len(entradas) == 0 {
		return ""
	}
	if frase, err := entradas[0].FormatFirst(); err == nil && frase != "" {
		return frase
	}
	if valor, err := entradas[0].Value(); err == nil {
		return fmt.Sprint(valor)
	}
	return ""
}

