package contract

import (
	"context"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

type WeatherService interface {
	GetWeather(ctx context.Context, location string) (domain.WeatherResponse, error)
}
