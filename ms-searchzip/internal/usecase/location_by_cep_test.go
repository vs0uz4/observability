package usecase

import (
	"errors"
	"testing"

	"github.com/vs0uz4/observability/ms-searchzip/internal/domain"
	"github.com/vs0uz4/observability/ms-searchzip/internal/service/mock"
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
		result := isNumeric(tt.input)
		if result != tt.expected {
			t.Errorf("For input %q, expected %v, got %v", tt.input, tt.expected, result)
		}
	}
}

func TestNewLocationByCepUsecase(t *testing.T) {
	mockCepSvc := &mock.MockCepService{}

	usecase := NewLocationByCepUsecase(mockCepSvc)

	if usecase.CepService != mockCepSvc {
		t.Errorf("Expected CepService to be %v, got %v", mockCepSvc, usecase.CepService)
	}
}

func TestGetLocationByCep(t *testing.T) {
	tests := []struct {
		name         string
		inputCep     string
		mockCepSvc   func() *mock.MockCepService
		expectErr    error
		expectOutput domain.CepResponse
	}{
		{
			name:     "Invalid CEP",
			inputCep: "123",
			mockCepSvc: func() *mock.MockCepService {
				return &mock.MockCepService{}
			},
			expectErr: domain.ErrInvalidZipcode,
		},
		{
			name:     "CEP Not Found",
			inputCep: "99999999",
			mockCepSvc: func() *mock.MockCepService {
				return &mock.MockCepService{
					GetLocationFunc: func(cep string) (domain.CepResponse, error) {
						return domain.CepResponse{}, domain.ErrZipcodeNotFound
					},
				}
			},
			expectErr: domain.ErrZipcodeNotFound,
		},
		{
			name:     "CEP Service Error",
			inputCep: "12345678",
			mockCepSvc: func() *mock.MockCepService {
				return &mock.MockCepService{
					GetLocationFunc: func(cep string) (domain.CepResponse, error) {
						return domain.CepResponse{}, domain.ErrCepService
					},
				}
			},
			expectErr: domain.ErrCepService,
		},
		{
			name:     "Success",
			inputCep: "12345678",
			mockCepSvc: func() *mock.MockCepService {
				return &mock.MockCepService{
					GetLocationFunc: func(cep string) (domain.CepResponse, error) {
						return domain.CepResponse{
							Localidade: "City",
							Uf:         "State",
						}, nil
					},
				}
			},
			expectOutput: domain.CepResponse{
				Localidade: "City",
				Uf:         "State",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := locationByCepUsecase{
				CepService: tt.mockCepSvc(),
			}

			result, err := usecase.GetLocationByCep(tt.inputCep)

			if !errors.Is(err, tt.expectErr) {
				t.Errorf("Expected error %v, got %v", tt.expectErr, err)
			}

			if result != tt.expectOutput {
				t.Errorf("Expected output %+v, got %+v", tt.expectOutput, result)
			}
		})
	}
}
