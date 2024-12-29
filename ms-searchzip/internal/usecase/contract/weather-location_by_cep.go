package contract

import "github.com/vs0uz4/observability/ms-searchzip/internal/domain"

type WeatherLocationByCepUsecase interface {
	GetWeatherLocationByCep(cep string) (domain.WeatherResponse, error)
}
