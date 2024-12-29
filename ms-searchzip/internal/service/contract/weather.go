package contract

import "github.com/vs0uz4/observability/ms-searchzip/internal/domain"

type WeatherService interface {
	GetWeather(cep string) (domain.WeatherResponse, error)
}
