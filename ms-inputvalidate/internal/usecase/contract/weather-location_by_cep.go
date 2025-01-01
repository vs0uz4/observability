package contract

import (
	"context"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
)

type WeatherLocationByCepUsecase interface {
	GetWeatherLocationByCep(ctx context.Context, cep string) (domain.WeatherResponse, error)
}
