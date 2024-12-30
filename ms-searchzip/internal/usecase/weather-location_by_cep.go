package usecase

import (
	"github.com/vs0uz4/observability/ms-searchzip/internal/domain"
	"github.com/vs0uz4/observability/ms-searchzip/internal/infra/utils"
	"github.com/vs0uz4/observability/ms-searchzip/internal/service/contract"
)

type weatherLocationByCepUsecase struct {
	WeatherService contract.WeatherService
}

func NewWeatherLocationByCepUsecase(weatherService contract.WeatherService) *weatherLocationByCepUsecase {
	return &weatherLocationByCepUsecase{
		WeatherService: weatherService,
	}
}

func (uc *weatherLocationByCepUsecase) GetWeatherLocationByCep(cep string) (domain.WeatherResponse, error) {
	if len(cep) != 8 || !utils.IsNumeric(cep) {
		return domain.WeatherResponse{}, domain.ErrInvalidZipcode
	}

	weather, err := uc.WeatherService.GetWeather(cep)
	if err != nil {
		return domain.WeatherResponse{}, err
	}

	return weather, nil
}
