package mock

import (
	"context"
	"errors"
	"testing"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

func TestMockWeatherByCepUsecase(t *testing.T) {
	mockUsecase := &MockWeatherByCepUsecase{
		GetWeatherByCepFunc: func(ctx context.Context, cep string) (domain.WeatherResponse, error) {
			if cep == "12345678" {
				return domain.WeatherResponse{
					Current: domain.CurrentWeather{
						TempC: 25.5,
					},
				}, nil
			}
			return domain.WeatherResponse{}, errors.New("invalid cep")
		},
	}

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		resp, err := mockUsecase.GetWeatherByCep(ctx, "12345678")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if resp.Current.TempC != 25.5 {
			t.Errorf("Expected TempC 25.5, got %v", resp.Current.TempC)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		_, err := mockUsecase.GetWeatherByCep(ctx, "00000000")
		if err == nil || err.Error() != "invalid cep" {
			t.Errorf("Expected error 'invalid cep', got %v", err)
		}
	})
}
