package web

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/webserver/middleware"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/mock"
)

func TestGetLocationWeatherByCep(t *testing.T) {
	tests := []struct {
		name           string
		inputCEP       string
		mockUsecase    func() *mock.MockWeatherLocationByCepUsecase
		expectedStatus int
		expectedBody   string
		expectedError  string
	}{
		{
			name:     "Erro no JSON Encode",
			inputCEP: "12345678",
			mockUsecase: func() *mock.MockWeatherLocationByCepUsecase {
				return &mock.MockWeatherLocationByCepUsecase{
					GetWeatherLocationByCepFunc: func(cep string) (domain.WeatherResponse, error) {
						return domain.WeatherResponse{
							City:       "City",
							Celsius:    math.NaN(),
							Fahrenheit: math.NaN(),
							Kelvin:     math.NaN(),
						}, nil
					},
				}
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "internal server error",
		},
		{
			name:     "CEP Inválido",
			inputCEP: "123",
			mockUsecase: func() *mock.MockWeatherLocationByCepUsecase {
				return &mock.MockWeatherLocationByCepUsecase{
					GetWeatherLocationByCepFunc: func(cep string) (domain.WeatherResponse, error) {
						return domain.WeatherResponse{}, domain.ErrInvalidZipcode
					},
				}
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   "invalid zipcode",
			expectedError:  "Invalid zipcode",
		},
		{
			name:     "CEP Não Encontrado",
			inputCEP: "99999999",
			mockUsecase: func() *mock.MockWeatherLocationByCepUsecase {
				return &mock.MockWeatherLocationByCepUsecase{
					GetWeatherLocationByCepFunc: func(cep string) (domain.WeatherResponse, error) {
						return domain.WeatherResponse{}, domain.ErrZipcodeNotFound
					},
				}
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "can not find zipcode",
			expectedError:  "Zipcode not found",
		},
		{
			name:     "Erro no Serviço de CEP",
			inputCEP: "12345678",
			mockUsecase: func() *mock.MockWeatherLocationByCepUsecase {
				return &mock.MockWeatherLocationByCepUsecase{
					GetWeatherLocationByCepFunc: func(cep string) (domain.WeatherResponse, error) {
						return domain.WeatherResponse{}, domain.ErrCepService
					},
				}
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "internal server error",
			expectedError:  "Internal server error",
		},
		{
			name:     "Erro no Serviço de Clima",
			inputCEP: "12345678",
			mockUsecase: func() *mock.MockWeatherLocationByCepUsecase {
				return &mock.MockWeatherLocationByCepUsecase{
					GetWeatherLocationByCepFunc: func(cep string) (domain.WeatherResponse, error) {
						return domain.WeatherResponse{}, domain.ErrWeatherService
					},
				}
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "internal server error",
			expectedError:  "Internal server error",
		},
		{
			name:     "Sucesso",
			inputCEP: "12345678",
			mockUsecase: func() *mock.MockWeatherLocationByCepUsecase {
				return &mock.MockWeatherLocationByCepUsecase{
					GetWeatherLocationByCepFunc: func(cep string) (domain.WeatherResponse, error) {
						return domain.WeatherResponse{
							City:       "City",
							Celsius:    25,
							Fahrenheit: 77,
							Kelvin:     298.15,
						}, nil
					},
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"city":"City","temp_C":25,"temp_F":77,"temp_K":298.15}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := tt.mockUsecase()
			handler := NewInputHandler(mockUsecase)

			rr := &middleware.ResponseRecorder{ResponseWriter: httptest.NewRecorder()}

			requestBody := fmt.Sprintf(`{"cep": "%s"}`, tt.inputCEP)

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")

			handler.GetLocationWeatherByCep(rr, req)

			resp := rr.ResponseWriter.(*httptest.ResponseRecorder).Result()
			body := strings.TrimSpace(rr.ResponseWriter.(*httptest.ResponseRecorder).Body.String())

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if body != tt.expectedBody {
				t.Errorf("Expected body %q, got %q", tt.expectedBody, body)
			}

			if rr.ReadError() != tt.expectedError {
				t.Errorf("Expected WriteError %q, got %q", tt.expectedError, rr.ReadError())
			}
		})
	}
}

func TestNewInputHandlerInitialization(t *testing.T) {
	mockUsecase := &mock.MockWeatherLocationByCepUsecase{}
	handler := NewInputHandler(mockUsecase)

	if handler.Usecase != mockUsecase {
		t.Errorf("Expected usecase %v, got %v", mockUsecase, handler.Usecase)
	}
}

func TestGetLocationWeatherByCepInvalidJSON(t *testing.T) {
	invalidJSON := `{"cep": 12345678`

	mockUsecase := &mock.MockWeatherLocationByCepUsecase{}
	handler := NewInputHandler(mockUsecase)

	req := httptest.NewRequest(http.MethodPost, "/weather", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.GetLocationWeatherByCep(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	expectedBody := "invalid request\n"
	if string(body) != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, string(body))
	}
}
