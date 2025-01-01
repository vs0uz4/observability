package mock

import (
	"context"
	"testing"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
)

func TestMockWeatherLocationByCepUsecase(t *testing.T) {
	mockUsecase := &MockWeatherLocationByCepUsecase{
		GetWeatherLocationByCepFunc: func(ctx context.Context, cep string) (domain.WeatherResponse, error) {
			switch cep {
			case "12345678":
				return domain.WeatherResponse{
					City:       "ABC",
					Celsius:    28.2,
					Fahrenheit: 82.8,
					Kelvin:     301.33,
				}, nil
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
		resp, err := mockUsecase.GetWeatherLocationByCep(ctx, "12345678")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if resp.City != "ABC" {
			t.Errorf("Expected City: ABC, got %v", resp.City)
		}

		if resp.Celsius != 28.2 {
			t.Errorf("Expected Celsius: 28.2, got %v", resp.Celsius)
		}

		if resp.Fahrenheit != 82.8 {
			t.Errorf("Expected Fahrenheit: 82.8, got %v", resp.Fahrenheit)
		}

		if resp.Kelvin != 301.33 {
			t.Errorf("Expected Kelvin: 301.33, got %v", resp.Kelvin)
		}
	})

	t.Run("Invalid ZipCode Failure", func(t *testing.T) {
		_, err := mockUsecase.GetWeatherLocationByCep(ctx, "123")
		if err == nil || err.Error() != "invalid zipcode" {
			t.Errorf("Expected error 'invalid zipcode', got %v", err)
		}
	})

	t.Run("ZipCode Not Found Failure", func(t *testing.T) {
		_, err := mockUsecase.GetWeatherLocationByCep(ctx, "87654321")
		if err == nil || err.Error() != "zipcode not found" {
			t.Errorf("Expected error 'zipcode not found', got %v", err)
		}
	})

	t.Run("Unexpected Error Failure", func(t *testing.T) {
		_, err := mockUsecase.GetWeatherLocationByCep(ctx, "00000000")
		if err == nil || err.Error() != "internal server error" {
			t.Errorf("Expected error 'internal server error', got %v", err)
		}
	})
}
