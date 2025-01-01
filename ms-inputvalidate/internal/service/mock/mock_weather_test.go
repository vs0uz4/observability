package mock

import (
	"context"
	"testing"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
)

func TestMockWeatherService(t *testing.T) {
	mock := MockWeatherService{
		GetWeatherFunc: func(ctx context.Context, cep string) (domain.WeatherResponse, error) {
			switch cep {
			case "12345678":
				return domain.WeatherResponse{City: "ABC", Celsius: 24.2, Fahrenheit: 75.6, Kelvin: 297.33}, nil
			case "123":
				return domain.WeatherResponse{}, domain.ErrInvalidZipcode
			case "87654321":
				return domain.WeatherResponse{}, domain.ErrZipcodeNotFound
			default:
				return domain.WeatherResponse{}, domain.ErrInternalServer
			}
		},
	}

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		response, err := mock.GetWeather(ctx, "12345678")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if response.Celsius != 24.2 || response.City != "ABC" {
			t.Errorf("Expected Celsius: 24.2 and City: ABC, got: %v and %v", response.Celsius, response.City)
		}
	})

	t.Run("Invalid ZipCode Failure", func(t *testing.T) {
		_, err := mock.GetWeather(ctx, "123")
		if err != domain.ErrInvalidZipcode {
			t.Errorf("Expected error: %v, got: %v", domain.ErrInvalidZipcode, err)
		}
	})

	t.Run("ZipCode Not Found Failure", func(t *testing.T) {
		_, err := mock.GetWeather(ctx, "87654321")
		if err != domain.ErrZipcodeNotFound {
			t.Errorf("Expected error: %v, got: %v", domain.ErrZipcodeNotFound, err)
		}
	})

	t.Run("Unexpected Error Failure", func(t *testing.T) {
		_, err := mock.GetWeather(ctx, "23456789")
		if err != domain.ErrInternalServer {
			t.Errorf("Expected error: %v, got: %v", domain.ErrInternalServer, err)
		}
	})
}
