package main

import (
	"os"
	"path/filepath"
	"sort"
)

const (
	CarpetaPorDefecto = "media"
	ArchivoSalida     = "metadata.json"
)

// Escaner recorre una carpeta y extrae metadatos de los archivos de medio.
type Escaner struct {
	ultimaCarpeta string
}

func NuevoEscaner() *Escaner {
	return &Escaner{}
}

// Escanear lee todos los archivos de imagen y video de la carpeta.
func (e *Escaner) Escanear(carpeta string) (*Report, error) {
	if carpeta == "" {
		carpeta = CarpetaPorDefecto
	}
	e.ultimaCarpeta = carpeta

	rep := NuevoReport()

	entradas, err := os.ReadDir(carpeta)
	if err != nil {
		return nil, err
	}

	var rutas []string
	for _, ent := range entradas {
		if ent.IsDir() {
			continue
		}
		rutas = append(rutas, filepath.Join(carpeta, ent.Name()))
	}
	sort.Strings(rutas)

	for _, ruta := range rutas {
		if TipoDeArchivo(ruta) == "desconocido" {
			continue
		}
		fm, err := ExtraerExif(ruta)
		if err != nil {
			// Saltamos archivos que no se pueden leer sin detener el escaneo.
			continue
		}
		rep.Archivos = append(rep.Archivos, *fm)
		rep.Total++
	}

	return rep, nil
}

// Retry vuelve a ejecutar el ultimo escaneo (usa la carpeta del scan previo).
func (e *Escaner) Retry() (*Report, error) {
	carpeta := e.ultimaCarpeta
	if carpeta == "" {
		carpeta = CarpetaPorDefecto
	}
	return e.Escanear(carpeta)
}
