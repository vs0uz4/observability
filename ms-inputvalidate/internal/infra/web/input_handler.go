package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/middleware"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/contract"
)

type CepRequest struct {
	Cep string `json:"cep"`
}

type InputHandler struct {
	Usecase contract.WeatherLocationByCepUsecase
}

func NewInputHandler(uc contract.WeatherLocationByCepUsecase) *InputHandler {
	return &InputHandler{Usecase: uc}
}

func (h *InputHandler) GetLocationWeatherByCep(w http.ResponseWriter, r *http.Request) {
	var req CepRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	weather, err := h.Usecase.GetWeatherLocationByCep(req.Cep)
	if err != nil {
		if errors.Is(err, domain.ErrZipcodeNotFound) {
			if rr, ok := w.(*middleware.ResponseRecorder); ok {
				rr.WriteError("Zipcode not found")
			}
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}

		if err.Error() == "invalid zipcode" {
			if rr, ok := w.(*middleware.ResponseRecorder); ok {
				rr.WriteError("Invalid zipcode")
			}
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if rr, ok := w.(*middleware.ResponseRecorder); ok {
			rr.WriteError("Internal server error")
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"city":   weather.City,
		"temp_C": weather.Celsius,
		"temp_F": weather.Fahrenheit,
		"temp_K": weather.Kelvin,
	}); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
