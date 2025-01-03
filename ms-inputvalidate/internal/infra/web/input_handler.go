package web

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/middleware"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/contract"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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

// GetLocationWeatherByCep é o handler que valida o INPUT(CEP) e encaminha a consulta ao serviço de clima
// @Summary Valida e encaminha um INPUT(CEP) para consulta de clima
// @Description Valida se o INPUT(CEP) contém 8 dígitos e encaminha para o serviço de clima se for válido.
// @Tags Input
// @Accept  json
// @Produce  json
// @Param   cep  body CepRequest   true  "CEP"
// @Success 200 {object} domain.WeatherResponse
// @Failure 422 {string} string "invalid zipcode"
// @Failure 404 {string} string "can not find zipcode"
// @Failure 500 {string} string "internal server error"
// @Router / [post]
func (h *InputHandler) GetLocationWeatherByCep(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("ms-inputvalidate")
	ctx, span := tracer.Start(context.Background(), "inputvalidate-request")
	defer span.End()

	var req CepRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	span.SetAttributes(attribute.String("input.cep", req.Cep))

	weather, err := h.Usecase.GetWeatherLocationByCep(ctx, req.Cep)
	if err != nil {
		if errors.Is(err, domain.ErrZipcodeNotFound) {
			if rr, ok := w.(*middleware.ResponseRecorder); ok {
				rr.WriteError("Zipcode not found")
			}
			span.RecordError(err)
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}

		if err.Error() == "invalid zipcode" {
			if rr, ok := w.(*middleware.ResponseRecorder); ok {
				rr.WriteError("Invalid zipcode")
			}
			span.RecordError(err)
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if rr, ok := w.(*middleware.ResponseRecorder); ok {
			rr.WriteError("Internal server error")
		}
		span.RecordError(err)
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
		span.RecordError(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
