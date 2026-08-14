package observer

import (
	"observer-weather/internal/domain"

	"github.com/pterm/pterm"
)

type Statistics struct {
	maxTemp, minTemp, sumTemp float64
	count                     int
}

func NewStatistics() *Statistics {
	return &Statistics{minTemp: 1000.0}
}

func (s *Statistics) Update(data domain.WeatherData) {
	s.count++
	s.sumTemp += data.Temperature
	if data.Temperature > s.maxTemp {
		s.maxTemp = data.Temperature
	}
	if data.Temperature < s.minTemp {
		s.minTemp = data.Temperature
	}

	promedio := s.sumTemp / float64(s.count)

	pterm.FgLightMagenta.Printfln("[STATS] Promedio: %.1f°C | Máx: %.1f°C | Mín: %.1f°C",
		promedio, s.maxTemp, s.minTemp)
}
