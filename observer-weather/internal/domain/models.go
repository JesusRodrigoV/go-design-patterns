package domain

type WeatherData struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
}

type Observer interface {
	Update(data WeatherData)
}

type Subject interface {
	Register(o Observer)
	Remove(o Observer)
	Notify()
}
