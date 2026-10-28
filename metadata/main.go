package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	lector := bufio.NewReader(os.Stdin)
	escaner := NuevoEscaner()

	fmt.Println("Metadata Reader - consola")
	fmt.Println("Comandos: scan [carpeta]  |  retry  |  exit")
	fmt.Println("La carpeta por defecto es ./media")
	fmt.Println()

	for {
		fmt.Print("> ")
		linea, err := lector.ReadString('\n')
		if err != nil {
			fmt.Println()
			break
		}
		linea = strings.TrimSpace(linea)
		if linea == "" {
			continue
		}

		partes := strings.Fields(linea)
		cmd := strings.ToLower(partes[0])

		switch cmd {
		case "exit":
			fmt.Println("Cerrando la aplicacion.")
			return
		case "scan":
			carpeta := ""
			if len(partes) > 1 {
				carpeta = partes[1]
			}
			ejecutarScan(escaner, carpeta)
		case "retry":
			ejecutarRetry(escaner)
		default:
			fmt.Println("Comando no reconocido. Usa: scan [carpeta] | retry | exit")
		}
	}
}

func ejecutarScan(e *Escaner, carpeta string) {
	rep, err := e.Escanear(carpeta)
	if err != nil {
		fmt.Println("Error al escanear:", err)
		return
	}

	if !rep.HayInfo() {
		fmt.Println("No se encontro informacion de metadatos. No se genero el JSON.")
		return
	}

	if err := guardarJSON(rep); err != nil {
		fmt.Println("Error al guardar el JSON:", err)
		return
	}

	fmt.Printf("Procesados %d archivo(s). Metadatos guardados en %s\n", rep.Total, ArchivoSalida)
	mostrarResumen(rep)
}

func ejecutarRetry(e *Escaner) {
	rep, err := e.Retry()
	if err != nil {
		fmt.Println("Error al reintentar:", err)
		return
	}

	if !rep.HayInfo() {
		fmt.Println("No se encontro informacion de metadatos. No se genero el JSON.")
		return
	}

	if err := guardarJSON(rep); err != nil {
		fmt.Println("Error al guardar el JSON:", err)
		return
	}

	fmt.Printf("Reintento completado. %d archivo(s). Metadatos guardados en %s\n", rep.Total, ArchivoSalida)
	mostrarResumen(rep)
}

func guardarJSON(rep *Report) error {
	datos, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ArchivoSalida, datos, 0o644)
}

func mostrarResumen(rep *Report) {
	for _, fm := range rep.Archivos {
		fmt.Printf(" - %s [%s]", fm.Archivo, fm.Tipo)
		if fm.Dispositivo != nil {
			fmt.Printf(" dispositivo=%s %s", fm.Dispositivo.Marca, fm.Dispositivo.Modelo)
		}
		if fm.FechaHora != "" {
			fmt.Printf(" fecha=%s", fm.FechaHora)
		}
		if fm.Ubicacion != nil {
			fmt.Printf(" ubicacion=%.6f,%.6f", fm.Ubicacion.Lat, fm.Ubicacion.Lon)
		}
		if fm.Usuario != "" {
			fmt.Printf(" usuario=%s", fm.Usuario)
		}
		fmt.Println()
	}
}
