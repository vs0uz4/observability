package contract

import "github.com/vs0uz4/observability/ms-weatherzip/internal/domain"

type WeatherService interface {
	GetWeather(location string) (domain.WeatherResponse, error)
}
