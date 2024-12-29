package usecase

import (
	"github.com/vs0uz4/observability/ms-searchzip/internal/domain"
	"github.com/vs0uz4/observability/ms-searchzip/internal/infra/utils"
	"github.com/vs0uz4/observability/ms-searchzip/internal/service/contract"
)

type weatherLocationByCepUsecase struct {
	CepService     contract.CepService
	WeatherService contract.WeatherService
}

func NewWeatherLocationByCepUsecase(cepService contract.CepService, weatherService contract.WeatherService) *weatherLocationByCepUsecase {
	return &weatherLocationByCepUsecase{
		CepService:     cepService,
		WeatherService: weatherService,
	}
}

func (uc *weatherLocationByCepUsecase) GetWeatherLocationByCep(cep string) (domain.WeatherResponse, error) {
	if len(cep) != 8 || !utils.IsNumeric(cep) {
		return domain.WeatherResponse{}, domain.ErrInvalidZipcode
	}

	_, err := uc.CepService.GetLocation(cep)
	if err != nil {
		return domain.WeatherResponse{}, err
	}

	weather, err := uc.WeatherService.GetWeather(cep)
	if err != nil {
		return domain.WeatherResponse{}, err
	}

	return weather, nil
}
