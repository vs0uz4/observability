package usecase

import (
	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/service/contract"
)

type weatherByCepUsecase struct {
	CepService     contract.CepService
	WeatherService contract.WeatherService
}

func NewWeatherByCepUsecase(cepService contract.CepService, weatherService contract.WeatherService) *weatherByCepUsecase {
	return &weatherByCepUsecase{
		CepService:     cepService,
		WeatherService: weatherService,
	}
}

func (uc *weatherByCepUsecase) GetWeatherByCep(cep string) (domain.WeatherResponse, error) {
	if len(cep) != 8 || !isNumeric(cep) {
		return domain.WeatherResponse{}, domain.ErrInvalidZipcode
	}

	location, err := uc.CepService.GetLocation(cep)
	if err != nil {
		return domain.WeatherResponse{}, err
	}

	weather, err := uc.WeatherService.GetWeather(location.Localidade)
	if err != nil {
		return domain.WeatherResponse{}, err
	}

	return weather, nil
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
