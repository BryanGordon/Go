package main

import (
	"math"
	"os"
	"path/filepath"
	"testing"
)

func putU16LE(b *[]byte, v uint16) {
	*b = append(*b, byte(v), byte(v>>8))
}

func putU32LE(b *[]byte, v uint32) {
	*b = append(*b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

// buildExifTiff construye un bloque TIFF/EXIF minimo (little endian) con:
// Make, Modelo, DateTime, DateTimeOriginal, Artist y GPS (lat/lon/alt).
func buildExifTiff() []byte {
	type ent struct {
		tag   uint16
		typ   uint16 // 1=BYTE, 2=ASCII, 4=LONG, 5=RATIONAL
		count uint32
		val   uint32
	}
	var ifd0, gps []ent
	var blob []byte

	putAscii := func(dst *[]ent, tag uint16, s string) {
		data := append([]byte(s), 0)
		off := uint32(len(blob))
		blob = append(blob, data...)
		*dst = append(*dst, ent{tag, 2, uint32(len(data)), off})
	}
	putRat := func(dst *[]ent, tag uint16, num, den uint32) {
		off := uint32(len(blob))
		putU32LE(&blob, num)
		putU32LE(&blob, den)
		*dst = append(*dst, ent{tag, 5, 1, off})
	}
	putRat3 := func(dst *[]ent, tag uint16, t [3]struct{ n, d uint32 }) {
		off := uint32(len(blob))
		for i := range t {
			putU32LE(&blob, t[i].n)
			putU32LE(&blob, t[i].d)
		}
		*dst = append(*dst, ent{tag, 5, 3, off})
	}
	putByte := func(dst *[]ent, tag uint16, v byte) {
		*dst = append(*dst, ent{tag, 1, 1, uint32(v)})
	}

	putAscii(&ifd0, 0x010F, "Canon")       // Make
	putAscii(&ifd0, 0x0110, "EOS 5D")      // Model
	putAscii(&ifd0, 0x0132, "2024:01:15 10:30:00") // DateTime
	putAscii(&ifd0, 0x9003, "2024:01:15 10:30:00") // DateTimeOriginal
	putAscii(&ifd0, 0x013B, "juan")        // Artist (usuario)

	putAscii(&gps, 0x0001, "N")                                  // LatitudeRef
	putRat3(&gps, 0x0002, [3]struct{ n, d uint32 }{{40, 1}, {26, 1}, {468, 10}}) // Latitude
	putAscii(&gps, 0x0003, "W")                                  // LongitudeRef
	putRat3(&gps, 0x0004, [3]struct{ n, d uint32 }{{73, 1}, {59, 1}, {723, 10}}) // Longitude
	putByte(&gps, 0x0005, 0)                                     // AltitudeRef
	putRat(&gps, 0x0006, 50, 1)                                  // Altitude

	blobLen := uint32(len(blob))
	gpsIfdOffset := uint32(8+2+(len(ifd0)+1)*12+4) + blobLen
	ifd0 = append(ifd0, ent{0x8825, 4, 1, gpsIfdOffset}) // GPS IFD offset

	var out []byte
	out = append(out, "II"...)
	putU16LE(&out, 42)
	putU32LE(&out, 8) // first IFD offset

	putU16LE(&out, uint16(len(ifd0)))
	for _, e := range ifd0 {
		putU16LE(&out, e.tag)
		putU16LE(&out, e.typ)
		putU32LE(&out, e.count)
		putU32LE(&out, e.val)
	}
	putU32LE(&out, 0) // next IFD offset
	out = append(out, blob...)

	putU16LE(&out, uint16(len(gps)))
	for _, e := range gps {
		putU16LE(&out, e.tag)
		putU16LE(&out, e.typ)
		putU32LE(&out, e.count)
		putU32LE(&out, e.val)
	}
	putU32LE(&out, 0)
	return out
}

// buildJpeg envuelve un bloque EXIF dentro de un contenedor JPEG minimo valido.
func buildJpeg(exif []byte) []byte {
	var j []byte
	j = append(j, 0xFF, 0xD8) // SOI
	j = append(j, 0xFF, 0xE1) // APP1
	app1 := append([]byte("Exif\x00\x00"), exif...)
	putU16LE(&j, uint16(2+len(app1)))
	j = append(j, app1...)
	j = append(j, 0xFF, 0xD9) // EOI
	return j
}

func escribirArchivoTemp(t *testing.T, nombre string, datos []byte) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, nombre)
	if err := os.WriteFile(path, datos, 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func approx(a, b float64) bool {
	return math.Abs(a-b) < 0.001
}

func TestExtraerExifConDatosCompletos(t *testing.T) {
	path := escribirArchivoTemp(t, "foto.jpg", buildJpeg(buildExifTiff()))

	fm, err := ExtraerExif(path)
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}

	if fm.Archivo != "foto.jpg" {
		t.Errorf("archivo = %q, esperado %q", fm.Archivo, "foto.jpg")
	}
	if fm.Tipo != "imagen" {
		t.Errorf("tipo = %q, esperado %q", fm.Tipo, "imagen")
	}
	if fm.Dispositivo == nil {
		t.Fatal("dispositivo no debe ser nil")
	}
	if fm.Dispositivo.Marca != "Canon" {
		t.Errorf("marca = %q, esperado %q", fm.Dispositivo.Marca, "Canon")
	}
	if fm.Dispositivo.Modelo != "EOS 5D" {
		t.Errorf("modelo = %q, esperado %q", fm.Dispositivo.Modelo, "EOS 5D")
	}
	if fm.FechaHora != "2024:01:15 10:30:00" {
		t.Errorf("fecha_hora = %q, esperado %q", fm.FechaHora, "2024:01:15 10:30:00")
	}
	if fm.Usuario != "juan" {
		t.Errorf("usuario = %q, esperado %q", fm.Usuario, "juan")
	}
	if fm.Ubicacion == nil {
		t.Fatal("ubicacion no debe ser nil")
	}
	if !approx(fm.Ubicacion.Lat, 40+26.0/60+46.8/3600) {
		t.Errorf("lat = %f, esperado ~%f", fm.Ubicacion.Lat, 40+26.0/60+46.8/3600)
	}
	if !approx(fm.Ubicacion.Lon, -(73+59.0/60+72.3/3600)) {
		t.Errorf("lon = %f, esperado ~%f", fm.Ubicacion.Lon, -(73+59.0/60+72.3/3600))
	}
	if fm.Ubicacion.Alt == nil || *fm.Ubicacion.Alt != 50 {
		t.Errorf("altura = %v, esperado 50", fm.Ubicacion.Alt)
	}
}

func TestExtraerExifSinExifNoEsError(t *testing.T) {
	path := escribirArchivoTemp(t, "vacio.jpg", buildJpeg(nil))

	fm, err := ExtraerExif(path)
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if fm.Dispositivo != nil || fm.Ubicacion != nil || fm.FechaHora != "" || fm.Usuario != "" {
		t.Error("no se esperaba metadato en archivo sin EXIF")
	}
}

func TestReportHayInfo(t *testing.T) {
	rep := NuevoReport()
	if rep.HayInfo() {
		t.Error("reporte vacio no debe tener informacion")
	}
	rep.Archivos = append(rep.Archivos, FileMeta{Archivo: "a.jpg", Tipo: "imagen"})
	rep.Total = 1
	if !rep.HayInfo() {
		t.Error("reporte con un archivo debe tener informacion")
	}
}
