package entity

import "regexp"

type CepTemperature struct {
	Cep       string
	Location  string
	Farenheit float64
	Celsius   float64
	Kelvin    float64
}

func newCepTemperature(cep string, location string, farenheit, celsius, kelvin float64) (*CepTemperature, error) {
	return &CepTemperature{
		Cep:       cep,
		Location:  location,
		Farenheit: farenheit,
		Celsius:   celsius,
		Kelvin:    kelvin,
	}, nil

}

func (c *CepTemperature) IsValid() bool {
	// CEP format: 99.999-999
	re := regexp.MustCompile(`^\d{2}\.\d{3}-\d{3}$`)
	if !re.MatchString(c.Cep) || c.Cep == "" {
		return false
	}

	return true
}

func (c *CepTemperature) ConvertTemperature() {
	c.Farenheit = celsiusToFahrenheit(c.Celsius)
	c.Kelvin = celsiusToKelvin(c.Celsius)
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}
