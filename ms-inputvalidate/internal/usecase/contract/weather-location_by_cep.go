package contract

import "github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"

type WeatherLocationByCepUsecase interface {
	GetWeatherLocationByCep(cep string) (domain.WeatherResponse, error)
}
