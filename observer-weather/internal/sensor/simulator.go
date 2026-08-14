package sensor

import (
	"math/rand"
	"observer-weather/internal/station"
	"time"
)

func StartSensors(ws *station.WeatherStation) {
	go simulateTemperature(ws)
	go simulateHumidity(ws)
	go simulatePressure(ws)
}

func simulateTemperature(ws *station.WeatherStation) {
	for {
		t := 20.0 + rand.Float64()*(40.0-20.0)
		ws.UpdateTemperature(t)
		time.Sleep(2 * time.Second)
	}
}

func simulateHumidity(ws *station.WeatherStation) {
	for {
		h := 40.0 + rand.Float64()*(95.0-40.0)
		ws.UpdateHumidity(h)
		time.Sleep(3 * time.Second)
	}
}

func simulatePressure(ws *station.WeatherStation) {
	for {
		p := 980.0 + rand.Float64()*(1030.0-980.0)
		ws.UpdatePressure(p)
		time.Sleep(4 * time.Second)
	}
}
