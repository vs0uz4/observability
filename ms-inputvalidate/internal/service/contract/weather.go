package contract

import "github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"

type WeatherService interface {
	GetWeather(cep string) (domain.WeatherResponse, error)
}
