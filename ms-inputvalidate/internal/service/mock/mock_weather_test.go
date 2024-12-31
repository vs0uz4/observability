package mock

import (
	"testing"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
)

func TestMockWeatherService(t *testing.T) {
	mock := MockWeatherService{
		GetWeatherFunc: func(cep string) (domain.WeatherResponse, error) {
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

	response, err := mock.GetWeather("12345678")
	if response.Celsius != 24.2 && response.City != "ABC" || err != nil {
		t.Errorf("Expected Celsius: 24.2 and City: ABC, got: %v, err: %v", response.Celsius, err)
	}

	expectedMessage := "Expected error: %v, got: %v"
	_, err = mock.GetWeather("123")
	if err != domain.ErrInvalidZipcode {
		t.Errorf(expectedMessage, domain.ErrInvalidZipcode, err)
	}

	_, err = mock.GetWeather("87654321")
	if err != domain.ErrZipcodeNotFound {
		t.Errorf(expectedMessage, domain.ErrZipcodeNotFound, err)
	}

	_, err = mock.GetWeather("23456789")
	if err != domain.ErrInternalServer {
		t.Errorf(expectedMessage, domain.ErrInternalServer, err)
	}
}
