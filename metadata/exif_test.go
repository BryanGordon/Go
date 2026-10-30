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
// En TIFF, los valores ASCII/Rational con count > 4 van en un offset externo;
// los de count <= 4 se empaquetan inline en los 4 bytes del campo val.
func buildExifTiff() []byte {
	type ent struct {
		tag   uint16
		typ   uint16 // 1=BYTE, 2=ASCII, 4=LONG, 5=RATIONAL
		count uint32
		val   uint32 // offset absoluto o valor inline segun el caso
	}
	var ifd0, gps []ent
	var blobIFD0, blobGPS []byte

	const (
		tiffHeaderSize = 8
		ifdEntrySize   = 12
	)

	// putAscii guarda s (con null) en blob y devuelve su offset relativo.
	putAscii := func(blob *[]byte, s string) uint32 {
		data := append([]byte(s), 0)
		off := uint32(len(*blob))
		*blob = append(*blob, data...)
		return off
	}
	// asciiInline empaqueta s (con null, hasta 4 bytes) en un uint32 little-endian.
	asciiInline := func(s string) uint32 {
		b := append([]byte(s), 0)
		for len(b) < 4 {
			b = append(b, 0)
		}
		return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
	}
	putRat := func(blob *[]byte, num, den uint32) uint32 {
		off := uint32(len(*blob))
		putU32LE(blob, num)
		putU32LE(blob, den)
		return off
	}
	putRat3 := func(blob *[]byte, t [3]struct{ n, d uint32 }) uint32 {
		off := uint32(len(*blob))
		for i := range t {
			putU32LE(blob, t[i].n)
			putU32LE(blob, t[i].d)
		}
		return off
	}

	// Datos de valores de IFD0
	offMake := putAscii(&blobIFD0, "Canon")
	offModel := putAscii(&blobIFD0, "EOS 5D")
	offDateTime := putAscii(&blobIFD0, "2024:01:15 10:30:00")
	offDateTimeOrig := putAscii(&blobIFD0, "2024:01:15 10:30:00")
	offArtist := putAscii(&blobIFD0, "juan")

	// Datos de valores de GPS
	offLat := putRat3(&blobGPS, [3]struct{ n, d uint32 }{{40, 1}, {26, 1}, {468, 10}})
	offLon := putRat3(&blobGPS, [3]struct{ n, d uint32 }{{73, 1}, {59, 1}, {723, 10}})
	blobGPS = append(blobGPS, 0) // AltitudeRef = 0
	offAlt := putRat(&blobGPS, 50, 1)

	numEntriesIFD0 := 6
	numEntriesGPS := 6

	dataOffsetIFD0 := uint32(tiffHeaderSize) + 2 + uint32(numEntriesIFD0)*uint32(ifdEntrySize) + 4
	gpsIfdOffset := dataOffsetIFD0 + uint32(len(blobIFD0))
	dataOffsetGPS := gpsIfdOffset + 2 + uint32(numEntriesGPS)*uint32(ifdEntrySize) + 4

	abs := func(base uint32, rel uint32) uint32 { return base + rel }

	ifd0 = append(ifd0, ent{0x010F, 2, uint32(len("Canon")+1), abs(dataOffsetIFD0, offMake)})       // Make
	ifd0 = append(ifd0, ent{0x0110, 2, uint32(len("EOS 5D")+1), abs(dataOffsetIFD0, offModel)})      // Model
	ifd0 = append(ifd0, ent{0x0132, 2, uint32(len("2024:01:15 10:30:00")+1), abs(dataOffsetIFD0, offDateTime)})
	ifd0 = append(ifd0, ent{0x9003, 2, uint32(len("2024:01:15 10:30:00")+1), abs(dataOffsetIFD0, offDateTimeOrig)})
	ifd0 = append(ifd0, ent{0x013B, 2, uint32(len("juan")+1), abs(dataOffsetIFD0, offArtist)})        // Artist
	ifd0 = append(ifd0, ent{0x8825, 4, 1, gpsIfdOffset})                                                // GPS IFD offset

	gps = append(gps, ent{0x0001, 2, uint32(len("N")+1), asciiInline("N")})          // LatitudeRef (inline)
	gps = append(gps, ent{0x0002, 5, 3, abs(dataOffsetGPS, offLat)})                    // Latitude
	gps = append(gps, ent{0x0003, 2, uint32(len("W")+1), asciiInline("W")})          // LongitudeRef (inline)
	gps = append(gps, ent{0x0004, 5, 3, abs(dataOffsetGPS, offLon)})                    // Longitude
	gps = append(gps, ent{0x0005, 1, 1, 0})                                             // AltitudeRef (inline = 0)
	gps = append(gps, ent{0x0006, 5, 1, abs(dataOffsetGPS, offAlt)})                    // Altitude

	var out []byte
	out = append(out, "II"...) // little endian
	putU16LE(&out, 42)
	putU32LE(&out, tiffHeaderSize) // first IFD offset = 8

	putU16LE(&out, uint16(numEntriesIFD0))
	for _, e := range ifd0 {
		putU16LE(&out, e.tag)
		putU16LE(&out, e.typ)
		putU32LE(&out, e.count)
		putU32LE(&out, e.val)
	}
	putU32LE(&out, 0) // next IFD offset
	out = append(out, blobIFD0...)

	putU16LE(&out, uint16(numEntriesGPS))
	for _, e := range gps {
		putU16LE(&out, e.tag)
		putU16LE(&out, e.typ)
		putU32LE(&out, e.count)
		putU32LE(&out, e.val)
	}
	putU32LE(&out, 0)
	out = append(out, blobGPS...)

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
