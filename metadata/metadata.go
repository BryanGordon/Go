package main

import "time"

// FileMeta agrupa los metadatos de un solo archivo de imagen o video.
type FileMeta struct {
	Archivo   string     `json:"archivo"`
	Tipo      string     `json:"tipo"` // imagen | video
	Ubicacion *Location  `json:"ubicacion,omitempty"`
	Dispositivo *Device   `json:"dispositivo,omitempty"`
	FechaHora string     `json:"fecha_hora,omitempty"`
	Usuario   string     `json:"usuario,omitempty"`
}

// Location representa una coordenada GPS.
type Location struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Alt  *float64 `json:"alt,omitempty"`
}

// Device representa la marca y modelo del dispositivo de captura.
type Device struct {
	Marca  string `json:"marca,omitempty"`
	Modelo string `json:"modelo,omitempty"`
}

// Report es el resultado de un escaneo listo para serializar a JSON.
type Report struct {
	Archivos []FileMeta `json:"archivos"`
	Generado string     `json:"generado"`
	Total    int        `json:"total"`
}

// HayInfo indica si se encontro al menos un archivo de medio con información.
// El JSON solo se genera cuando esto es verdadero.
func (rep *Report) HayInfo() bool {
	return rep.Total > 0
}

func NuevoReport() *Report {
	return &Report{Generado: time.Now().Format(time.RFC3339)}
}
