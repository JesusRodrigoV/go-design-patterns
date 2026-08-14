package main

import (
	"fmt"
	"observer-weather/internal/observer"
	"observer-weather/internal/sensor"
	"observer-weather/internal/station"
	"time"
)

func main() {
	fmt.Println("Iniciando Sistema Meteorológico...")

	centralStation := station.New()

	display := &observer.Display{}
	stats := observer.NewStatistics()
	alert := &observer.Alert{}
	logger := &observer.Logger{}

	centralStation.Register(display)
	centralStation.Register(stats)
	centralStation.Register(alert)
	centralStation.Register(logger)

	sensor.StartSensors(centralStation)

	time.Sleep(10 * time.Second)
	fmt.Println("Ejecución finalizada.")
}
