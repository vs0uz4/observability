package mock

import (
	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

type MockWeatherByCepUsecase struct {
	GetWeatherByCepFunc func(cep string) (domain.WeatherResponse, error)
}

func (m *MockWeatherByCepUsecase) GetWeatherByCep(cep string) (domain.WeatherResponse, error) {
	return m.GetWeatherByCepFunc(cep)
}
