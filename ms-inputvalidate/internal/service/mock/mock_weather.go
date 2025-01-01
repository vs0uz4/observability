package mock

import (
	"context"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
)

type MockWeatherService struct {
	GetWeatherFunc func(ctx context.Context, cep string) (domain.WeatherResponse, error)
}

func (m *MockWeatherService) GetWeather(ctx context.Context, cep string) (domain.WeatherResponse, error) {
	return m.GetWeatherFunc(ctx, cep)
}
