package contract

import "github.com/vs0uz4/observability/ms-weatherzip/internal/domain"

type WeatherByCepUsecase interface {
	GetWeatherByCep(cep string) (domain.WeatherResponse, error)
}
