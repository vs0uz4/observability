package web

import (
	"encoding/json"
	"net/http"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/usecase"
)

type HealthHandler struct {
	useCase usecase.HealthCheckUseCase
}

func NewHealthHandler(u usecase.HealthCheckUseCase) *HealthHandler {
	return &HealthHandler{u}
}

// GetHealth é o handler responsável por retornar informações sobre a saúde do serviço
// @Summary Retorna informações sobre a saúde do serviço
// @Description Retorna informações sobre a saúde do serviço, como uso de CPU, memória e tempo de atividade
// @Tags Health
// @Produce  json
// @Success 200 {object} health.HealthStats
// @Failure 500 {string} string "Internal Server Error"
// @Router /health [get]
func (h *HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	health, err := h.useCase.GetHealth()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(health)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
