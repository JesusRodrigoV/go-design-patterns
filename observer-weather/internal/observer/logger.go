package observer

import (
	"observer-weather/internal/domain"

	"github.com/pterm/pterm"
)

type Logger struct{}

func (l *Logger) Update(data domain.WeatherData) {
	pterm.FgGray.Printfln("[LOG] Estado guardado -> T:%.1f H:%.1f P:%.1f",
		data.Temperature, data.Humidity, data.Pressure)
}
