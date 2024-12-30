package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/vs0uz4/observability/ms-searchzip/internal/domain"
	"github.com/vs0uz4/observability/ms-searchzip/internal/service/mock"
)

const (
	weatherZipServiceBaseURL = "http://example.com/weather/%s"
)

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated read error")
}

func (e *errorReader) Close() error {
	return nil
}

func TestWeatherServiceCreateRequest(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		expectErr string
	}{
		{
			name:      "Request Creation Error",
			inputURL:  "",
			expectErr: "failed to create request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WeatherService{}
			_, err := s.GetWeather(tt.inputURL)

			if err == nil || !strings.Contains(err.Error(), tt.expectErr) {
				t.Errorf("Expected error containing %q, got %v", tt.expectErr, err)
			}
		})
	}
}

func TestWeatherServiceExecuteRequest(t *testing.T) {
	tests := []struct {
		name      string
		expectErr string
	}{
		{
			name:      "Request Execution Error",
			expectErr: "failed to make request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mock.MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("network error")
				},
			}

			service := WeatherService{
				HttpClient: mockClient,
				BaseURL:    weatherZipServiceBaseURL,
			}

			_, err := service.GetWeather("12345678")

			if err == nil || !strings.Contains(err.Error(), tt.expectErr) {
				t.Errorf("Expected error containing %q, got %q", tt.expectErr, err.Error())
			}
		})
	}
}

func TestWeatherServiceDecodeResponse(t *testing.T) {
	tests := []struct {
		name          string
		inputResponse string
		expectErr     string
	}{
		{
			name:          "Failed to Decode Response",
			inputResponse: "invalid_json",
			expectErr:     "invalid character",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response domain.WeatherResponse
			err := json.Unmarshal([]byte(tt.inputResponse), &response)

			if err == nil || !strings.Contains(err.Error(), tt.expectErr) {
				t.Errorf("Expected error containing %q, got %q", tt.expectErr, err.Error())
			}
		})
	}
}

func TestWeatherServiceDecodeResponseError(t *testing.T) {
	mockClient := &mock.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`invalid_json`)),
			}, nil
		},
	}

	service := WeatherService{
		HttpClient: mockClient,
		BaseURL:    weatherZipServiceBaseURL,
	}

	_, err := service.GetWeather("12345678")

	if err == nil || !strings.Contains(err.Error(), "failed to decode response") {
		t.Errorf("Expected error containing %q, got %q", "failed to decode response", err.Error())
	}
}

func TestWeatherServiceReadResponseBodyError(t *testing.T) {
	mockClient := &mock.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusUnprocessableEntity,
				Body:       &errorReader{},
			}, nil
		},
	}

	service := WeatherService{
		HttpClient: mockClient,
		BaseURL:    weatherZipServiceBaseURL,
	}

	_, err := service.GetWeather("12345")

	if err == nil || !strings.Contains(err.Error(), "failed to read response body") {
		t.Errorf("Expected error containing %q, got %q", "failed to read response body", err.Error())
	}
}

func TestWeatherServiceUnexpectedErrors(t *testing.T) {
	tests := []struct {
		name      string
		inputCep  string
		inputCode int
		expectErr error
	}{
		{
			name:      "Unexpected Internal Server Error",
			inputCep:  "123",
			inputCode: 500,
			expectErr: domain.ErrUnexpectedInternalServer,
		},
		{
			name:      "Unexpected Unprocessable Entity Error",
			inputCep:  "456",
			inputCode: 422,
			expectErr: domain.ErrUnexpectedUnprocessableEntity,
		},
		{
			name:      "Unexpected Not Found Error",
			inputCep:  "789",
			inputCode: 404,
			expectErr: domain.ErrUnexpectedNotFound,
		},
		{
			name:      "Unexpected Weather Service Error",
			inputCep:  "000",
			inputCode: 400,
			expectErr: domain.NewUnexpectedWeatherServiceError(
				domain.ErrWeatherService,
				400,
				"",
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mock.MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					switch tt.inputCep {
					case "123":
						return &http.Response{
							StatusCode: http.StatusInternalServerError,
							Body:       io.NopCloser(strings.NewReader("unexpected internal server error")),
						}, nil
					case "456":
						return &http.Response{
							StatusCode: http.StatusUnprocessableEntity,
							Body:       io.NopCloser(strings.NewReader("unexpected unprocessable entity error")),
						}, nil
					case "789":
						return &http.Response{
							StatusCode: http.StatusNotFound,
							Body:       io.NopCloser(strings.NewReader("unexpected not found error")),
						}, nil
					default:
						return &http.Response{
							StatusCode: http.StatusBadRequest,
							Body:       io.NopCloser(strings.NewReader("")),
						}, nil
					}

				},
			}

			service := WeatherService{
				HttpClient: mockClient,
				BaseURL:    weatherZipServiceBaseURL,
			}
			_, err := service.GetWeather(tt.inputCep)

			if err == nil || err.Error() != tt.expectErr.Error() {
				t.Errorf("Expected error %v, got %v", tt.expectErr, err)
			}
		})
	}
}

func TestWeatherServiceGetWeatherData(t *testing.T) {
	tests := []struct {
		name           string
		mockResponse   string
		mockStatusCode int
		inputCep       string
		expectErr      error
		expectOutput   domain.WeatherResponse
	}{
		{
			name:           "12345678",
			mockResponse:   `{"City": "ABC", "temp_C": 25.0, "temp_F": 77.0, "temp_K": 313}`,
			mockStatusCode: http.StatusOK,
			inputCep:       "12345678",
			expectErr:      nil,
			expectOutput:   domain.WeatherResponse{City: "ABC", Celsius: 25.0, Fahrenheit: 77.0, Kelvin: 313},
		},
		{
			name:           "00000000",
			mockResponse:   "can not find zipcode",
			mockStatusCode: http.StatusNotFound,
			inputCep:       "00000000",
			expectErr:      domain.ErrZipcodeNotFound,
			expectOutput:   domain.WeatherResponse{},
		},
		{
			name:           "123",
			mockResponse:   "invalid zipcode",
			mockStatusCode: http.StatusUnprocessableEntity,
			inputCep:       "123",
			expectErr:      domain.ErrInvalidZipcode,
			expectOutput:   domain.WeatherResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				if _, err := w.Write([]byte(tt.mockResponse)); err != nil {
					t.Fatalf("Failed to write mock response: %v", err)
				}
			}))
			defer mockServer.Close()

			weatherZipService := NewWeatherService(mockServer.Client(), mockServer.URL+"/weather/%s")
			result, err := weatherZipService.GetWeather(tt.inputCep)

			if !errors.Is(err, tt.expectErr) {
				t.Errorf("Expected error %v, got %v", tt.expectErr, err)
			}

			if result != tt.expectOutput {
				t.Errorf("Expected output %+v, got %+v", tt.expectOutput, result)
			}
		})
	}
}
