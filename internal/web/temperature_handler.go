package web

import (
	"encoding/json"
	"net/http"

	"github.com/thiagohmm/fulcycleTemperaturaPorCep/internal/usecase"
)

type WeatherHandler struct {
	UseCase *usecase.TemperatureUseCase
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	var dto usecase.TemperatureInputDTO
	if len(cep) != 8 {
		http.Error(w, "Invalid CEP", http.StatusBadRequest)
		return
	}
	dto = usecase.TemperatureInputDTO{Cep: cep}

	weather, err := h.UseCase.Execute(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}
