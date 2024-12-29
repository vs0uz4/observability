package mock

import "github.com/vs0uz4/observability/ms-searchzip/internal/domain"

type MockWeatherService struct {
	GetWeatherFunc func(string) (domain.WeatherResponse, error)
}

func (m *MockWeatherService) GetWeather(cep string) (domain.WeatherResponse, error) {
	return m.GetWeatherFunc(cep)
}
