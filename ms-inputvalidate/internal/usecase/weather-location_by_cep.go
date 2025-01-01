package usecase

import (
	"context"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/utils"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/service/contract"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type weatherLocationByCepUsecase struct {
	WeatherService contract.WeatherService
}

func NewWeatherLocationByCepUsecase(weatherService contract.WeatherService) *weatherLocationByCepUsecase {
	return &weatherLocationByCepUsecase{
		WeatherService: weatherService,
	}
}

func (uc *weatherLocationByCepUsecase) GetWeatherLocationByCep(ctx context.Context, cep string) (domain.WeatherResponse, error) {
	tracer := otel.Tracer("ms-inputvalidate")
	ctx, span := tracer.Start(ctx, "GetWeatherLocationByCep")
	span.SetAttributes(attribute.String("cep", cep))
	defer span.End()

	if len(cep) != 8 || !utils.IsNumeric(cep) {
		span.RecordError(domain.NewInvalidZipCodeDetailsError(cep))
		return domain.WeatherResponse{}, domain.ErrInvalidZipcode
	}

	weather, err := uc.WeatherService.GetWeather(ctx, cep)
	if err != nil {
		span.RecordError(err)
		return domain.WeatherResponse{}, err
	}

	return weather, nil
}
