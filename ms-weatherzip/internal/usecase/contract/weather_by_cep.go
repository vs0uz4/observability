package contract

import (
	"context"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

type WeatherByCepUsecase interface {
	GetWeatherByCep(ctx context.Context, cep string) (domain.WeatherResponse, error)
}
