package observer

import (
	"observer-weather/internal/domain"

	"github.com/pterm/pterm"
)

type Display struct{}

func (d *Display) Update(data domain.WeatherData) {
	pterm.Info.Printfln("Temp: %.1f°C | Hum: %.1f%% | Pres: %.1f hPa",
		data.Temperature, data.Humidity, data.Pressure)
}
