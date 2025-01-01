package mock

import (
	"context"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

type MockWeatherService struct {
	GetWeatherFunc func(ctx context.Context, location string) (domain.WeatherResponse, error)
}

func (m *MockWeatherService) GetWeather(ctx context.Context, location string) (domain.WeatherResponse, error) {
	return m.GetWeatherFunc(ctx, location)
}
