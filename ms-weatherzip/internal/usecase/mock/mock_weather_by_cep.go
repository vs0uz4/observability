package mock

import (
	"context"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

type MockWeatherByCepUsecase struct {
	GetWeatherByCepFunc func(ctx context.Context, cep string) (domain.WeatherResponse, error)
}

func (m *MockWeatherByCepUsecase) GetWeatherByCep(ctx context.Context, cep string) (domain.WeatherResponse, error) {
	return m.GetWeatherByCepFunc(ctx, cep)
}
