package mock

import (
	"context"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
)

type MockWeatherLocationByCepUsecase struct {
	GetWeatherLocationByCepFunc func(ctx context.Context, cep string) (domain.WeatherResponse, error)
}

func (m *MockWeatherLocationByCepUsecase) GetWeatherLocationByCep(ctx context.Context, cep string) (domain.WeatherResponse, error) {
	return m.GetWeatherLocationByCepFunc(ctx, cep)
}
