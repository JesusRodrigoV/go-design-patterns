package station

import (
	"observer-weather/internal/domain"
	"sync"
)

type WeatherStation struct {
	mu        sync.RWMutex
	observers []domain.Observer
	state     domain.WeatherData
}

func New() *WeatherStation {
	return &WeatherStation{
		observers: make([]domain.Observer, 0),
	}
}

func (ws *WeatherStation) Register(o domain.Observer) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.observers = append(ws.observers, o)
}

func (ws *WeatherStation) Remove(o domain.Observer) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	for i, observer := range ws.observers {
		if observer == o {
			ws.observers = append(ws.observers[:i], ws.observers[i+1:]...)
			break
		}
	}
}

func (ws *WeatherStation) Notify() {
	ws.mu.RLock()
	currentState := ws.state
	currentObservers := make([]domain.Observer, len(ws.observers))
	copy(currentObservers, ws.observers)
	ws.mu.RUnlock()

	for _, observer := range currentObservers {
		observer.Update(currentState)
	}
}

func (ws *WeatherStation) UpdateTemperature(t float64) {
	ws.mu.Lock()
	ws.state.Temperature = t
	ws.mu.Unlock()
	ws.Notify()
}

func (ws *WeatherStation) UpdateHumidity(h float64) {
	ws.mu.Lock()
	ws.state.Humidity = h
	ws.mu.Unlock()
	ws.Notify()
}

func (ws *WeatherStation) UpdatePressure(p float64) {
	ws.mu.Lock()
	ws.state.Pressure = p
	ws.mu.Unlock()
	ws.Notify()
}
