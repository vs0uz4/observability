package mock

import "github.com/vs0uz4/observability/ms-weatherzip/internal/domain"

type MockWeatherService struct {
	GetWeatherFunc func(string) (domain.WeatherResponse, error)
}

func (m *MockWeatherService) GetWeather(location string) (domain.WeatherResponse, error) {
	return m.GetWeatherFunc(location)
}
