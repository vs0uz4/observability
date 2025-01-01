package usecase

import (
	"context"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/service/contract"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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

func (uc *weatherByCepUsecase) GetWeatherByCep(ctx context.Context, cep string) (domain.WeatherResponse, error) {
	tracer := otel.Tracer("ms-weatherzip")
	ctx, span := tracer.Start(ctx, "GetWeatherByCep")
	span.SetAttributes(attribute.String("cep", cep))
	defer span.End()

	if len(cep) != 8 || !isNumeric(cep) {
		span.RecordError(domain.ErrInvalidZipcode)
		return domain.WeatherResponse{}, domain.ErrInvalidZipcode
	}

	location, err := uc.CepService.GetLocation(ctx, cep)
	if err != nil {
		span.RecordError(err)
		return domain.WeatherResponse{}, err
	}

	weather, err := uc.WeatherService.GetWeather(ctx, location.Localidade)
	if err != nil {
		span.RecordError(err)
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
