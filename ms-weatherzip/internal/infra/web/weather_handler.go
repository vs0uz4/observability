package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/infra/web/webserver/middleware"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/usecase/contract"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"

	"github.com/go-chi/chi/v5"
)

type WeatherResponseSummary struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type WeatherHandler struct {
	Usecase contract.WeatherByCepUsecase
}

func NewWeatherHandler(uc contract.WeatherByCepUsecase) *WeatherHandler {
	return &WeatherHandler{Usecase: uc}
}

// GetWeatherByCep é o handler que consulta o (CEP), encontra a localidade e retorna o clima atual.
// @Summary Consulta o clima atual de uma localidade a partir do CEP
// @Description Consulta o clima atual de uma localidade a partir do CEP. Retorna a cidade e a temperatura atual em Celsius, Fahrenheit e Kelvin.
// @Tags Weather
// @Produce  json
// @param cep path string true "CEP para buscar a temperatura"
// @Success 200 {object} WeatherResponseSummary
// @Failure 422 {string} string "invalid zipcode"
// @Failure 404 {string} string "can not find zipcode"
// @Failure 500 {string} string "internal server error"
// @Router /weather/{cep} [get]
func (h *WeatherHandler) GetWeatherByCep(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(r.Header))

	for k, v := range r.Header {
		fmt.Printf("Header[%q] = %q\n", k, v)
	}

	tracer := otel.Tracer("ms-weatherzip")

	ctx, span := tracer.Start(ctx, "weatherzip-request")
	defer span.End()

	cep := chi.URLParam(r, "cep")
	span.SetAttributes(attribute.String("input.cep", cep))

	weather, err := h.Usecase.GetWeatherByCep(ctx, cep)
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
		"city":   weather.Location.Name,
		"temp_C": weather.Current.TempC,
		"temp_F": weather.Current.TempF,
		"temp_K": weather.Current.TempK,
	}); err != nil {
		span.RecordError(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
