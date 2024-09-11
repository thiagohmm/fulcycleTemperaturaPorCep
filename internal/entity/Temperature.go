package entity

type Temperature struct {
	Farenheit float64
	Celsius   float64
	Kelvin    float64
}

func NewTemperature(kelvin float64) (*Temperature, error) {
	return &Temperature{

		Farenheit: kelvinToFahrenheit(kelvin),
		Celsius:   kelvinToCelsius(kelvin),
		Kelvin:    kelvin,
	}, nil

}

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func kelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}
