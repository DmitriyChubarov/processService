package main

import "github.com/DmitriyChubarov/processing/internal/app"

const (
	serviceName = "processingService"
)

func main() {
	app.Run(serviceName)
}
