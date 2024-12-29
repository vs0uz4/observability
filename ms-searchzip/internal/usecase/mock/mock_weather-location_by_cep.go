package mock

import (
	"github.com/vs0uz4/observability/ms-searchzip/internal/domain"
)

type MockWeatherLocationByCepUsecase struct {
	GetWeatherLocationByCepFunc func(cep string) (domain.WeatherResponse, error)
}

func (m *MockWeatherLocationByCepUsecase) GetWeatherLocationByCep(cep string) (domain.WeatherResponse, error) {
	return m.GetWeatherLocationByCepFunc(cep)
}
