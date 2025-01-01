package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/utils"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/service/mock"
)

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "123456", expected: true},
		{input: "123abc", expected: false},
		{input: "", expected: true},
		{input: "123 456", expected: false},
		{input: "!@#$%", expected: false},
	}

	for _, tt := range tests {
		result := utils.IsNumeric(tt.input)
		if result != tt.expected {
			t.Errorf("For input %q, expected %v, got %v", tt.input, tt.expected, result)
		}
	}
}

func TestNewWeatherLocationByCepUsecase(t *testing.T) {
	mockWeatherSvc := &mock.MockWeatherService{}

	usecase := NewWeatherLocationByCepUsecase(mockWeatherSvc)

	if usecase.WeatherService != mockWeatherSvc {
		t.Errorf("Expected WeatherService to be %v, got %v", mockWeatherSvc, usecase.WeatherService)
	}
}

func TestGetWeatherLocationByCep(t *testing.T) {
	tests := []struct {
		name           string
		inputCep       string
		mockWeatherSvc func() *mock.MockWeatherService
		expectErr      error
		expectOutput   domain.WeatherResponse
	}{
		{
			name:     "Invalid CEP",
			inputCep: "123",
			mockWeatherSvc: func() *mock.MockWeatherService {
				return &mock.MockWeatherService{}
			},
			expectErr: domain.ErrInvalidZipcode,
		},
		{
			name:     "CEP Not Found",
			inputCep: "99999999",
			mockWeatherSvc: func() *mock.MockWeatherService {
				return &mock.MockWeatherService{
					GetWeatherFunc: func(ctx context.Context, cep string) (domain.WeatherResponse, error) {
						return domain.WeatherResponse{}, domain.ErrZipcodeNotFound
					},
				}
			},
			expectErr: domain.ErrZipcodeNotFound,
		},
		{
			name:     "WeahterZip Service Error",
			inputCep: "12345678",
			mockWeatherSvc: func() *mock.MockWeatherService {
				return &mock.MockWeatherService{
					GetWeatherFunc: func(ctx context.Context, cep string) (domain.WeatherResponse, error) {
						return domain.WeatherResponse{}, domain.ErrWeatherService
					},
				}
			},
			expectErr: domain.ErrWeatherService,
		},
		{
			name:     "Success",
			inputCep: "12345678",
			mockWeatherSvc: func() *mock.MockWeatherService {
				return &mock.MockWeatherService{
					GetWeatherFunc: func(ctx context.Context, cep string) (domain.WeatherResponse, error) {
						return domain.WeatherResponse{
							City:       "City",
							Celsius:    28.2,
							Fahrenheit: 82.8,
							Kelvin:     301.33,
						}, nil
					},
				}
			},
			expectOutput: domain.WeatherResponse{
				City:       "City",
				Celsius:    28.2,
				Fahrenheit: 82.8,
				Kelvin:     301.33,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			usecase := weatherLocationByCepUsecase{
				WeatherService: tt.mockWeatherSvc(),
			}

			result, err := usecase.GetWeatherLocationByCep(ctx, tt.inputCep)

			if !errors.Is(err, tt.expectErr) {
				t.Errorf("Expected error %v, got %v", tt.expectErr, err)
			}

			if result != tt.expectOutput {
				t.Errorf("Expected output %+v, got %+v", tt.expectOutput, result)
			}
		})
	}
}
