package observer

import (
	"observer-weather/internal/domain"

	"github.com/pterm/pterm"
)

type Alert struct{}

func (a *Alert) Update(data domain.WeatherData) {
	if data.Temperature > 35.0 {
		pterm.Error.Printfln("Calor extremo detectado: %.1f°C", data.Temperature)
	}
	if data.Pressure < 990.0 {
		pterm.Warning.Printfln("Baja presión detectada. Posible tormenta: %.1f hPa", data.Pressure)
	}
}
