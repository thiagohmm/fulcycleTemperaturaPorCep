package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

type GeoResponse struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func getCepData(cep string) (*ViaCepResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cepData ViaCepResponse
	if err := json.Unmarshal(body, &cepData); err != nil {
		fmt.Println("Erro ao decodificar resposta da API ViaCEP:", string(body))
		return nil, err
	}

	return &cepData, nil
}

func getGeoData(bairro, localidade string) (*GeoResponse, error) {
	bairro = url.QueryEscape(bairro)
	localidade = url.QueryEscape(localidade)
	fmt.Println(localidade, bairro)
	resp, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s,%s&appid=8f2dfe379acdba84dfa143c5648000e3", bairro, localidade))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geoData []GeoResponse
	if err := json.Unmarshal(body, &geoData); err != nil {
		fmt.Println("Erro ao decodificar resposta da API OpenWeatherMap (Geo):", string(body))
		return nil, err
	}

	if len(geoData) == 0 {
		return nil, fmt.Errorf("no geo data found")
	}

	return &geoData[0], nil
}

func getWeatherData(lat, lon float64) (*WeatherResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=8f2dfe379acdba84dfa143c5648000e3", lat, lon))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		fmt.Println("Erro ao decodificar resposta da API OpenWeatherMap (Weather):", string(body))
		return nil, err
	}

	return &weatherData, nil
}

func getTemperatureByCep(cep string) (float64, error) {
	cepData, err := getCepData(cep)
	fmt.Println(cepData)
	if err != nil {
		return 0, err
	}

	geoData, err := getGeoData(cepData.Bairro, cepData.Localidade)
	if err != nil {
		return 0, err
	}

	weatherData, err := getWeatherData(geoData.Lat, geoData.Lon)
	if err != nil {
		return 0, err
	}

	return weatherData.Main.Temp, nil
}

func main() {
	cep := "03142-001"
	temp, err := getTemperatureByCep(cep)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("A temperatura para o CEP %s é %.2fK\n", cep, temp)
}
