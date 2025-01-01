package mock

import (
	"context"
	"testing"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

func TestMockWeatherService(t *testing.T) {
	mock := MockWeatherService{
		GetWeatherFunc: func(ctx context.Context, location string) (domain.WeatherResponse, error) {
			if location == "Valid Location" {
				return domain.WeatherResponse{Current: domain.CurrentWeather{TempC: 25.0}}, nil
			}
			return domain.WeatherResponse{}, domain.ErrUnexpectedBadRequest
		},
	}

	ctx := context.Background()

	t.Run("Valid Location", func(t *testing.T) {
		response, err := mock.GetWeather(ctx, "Valid Location")
		if response.Current.TempC != 25.0 || err != nil {
			t.Errorf("Expected TempC: 25.0, got: %v, err: %v", response.Current.TempC, err)
		}
	})

	t.Run("Invalid Location", func(t *testing.T) {
		_, err := mock.GetWeather(ctx, "Invalid Location")
		if err != domain.ErrUnexpectedBadRequest {
			t.Errorf("Expected error: %v, got: %v", domain.ErrUnexpectedBadRequest, err)
		}
	})
}
