package usecase

import (
	"context"

	"github.com/thiagohmm/fulcycleTemperaturaPorCep/internal/entity"
	"github.com/thiagohmm/fulcycleTemperaturaPorCep/internal/infraestructure"
)

type TemperatureInputDTO struct {
	Cep string `json:"cep"`
}

type TemperatureOutputDTO struct {
	Celsius   float64 `json:"celsius"`
	Farenheit float64 `json:"farenheit"`
	Kelvin    float64 `json:"kelvin"`
}

type TemperatureUseCase struct {
	Apiclient   infraestructure.GetTemperatureForCep
	Temperature entity.Temperature
}

func NewTemperatureUseCase(apiClient infraestructure.GetTemperatureForCep) *TemperatureUseCase {
	return &TemperatureUseCase{
		Apiclient:   apiClient,
		Temperature: entity.Temperature{}, // Inicializa a estrutura vazia
	}
}

func (t *TemperatureUseCase) Execute(ctx context.Context, input TemperatureInputDTO) (*TemperatureOutputDTO, error) {
	wheatherK, err := t.Apiclient.GetTemperatureByCep(ctx, input.Cep)
	if err != nil {
		return nil, err
	}

	temperature, err := entity.NewTemperature(wheatherK)
	if err != nil {
		return nil, err
	}

	return &TemperatureOutputDTO{
		Celsius:   temperature.Celsius,
		Farenheit: temperature.Farenheit,
		Kelvin:    temperature.Kelvin,
	}, nil
}
