package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/service/contract"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var _ contract.WeatherService = (*WeatherService)(nil)

var weatherErrorCodes = map[int]error{
	1003: domain.ErrParameterNotProvided,
	1005: domain.ErrApiUrlIsInvalid,
	1006: domain.ErrLocationNotFound,
	9000: domain.ErrJsonBodyIsInvalid,
	9001: domain.ErrTooManyLocations,
	9999: domain.ErrInternalApplication,
}

type WeatherService struct {
	HttpClient contract.HttpClient
	BaseURL    string
	ApiKey     string
	Language   string
}

func NewWeatherService(client *http.Client, baseURL, apiKey, language string) *WeatherService {
	return &WeatherService{
		HttpClient: client,
		BaseURL:    baseURL,
		ApiKey:     apiKey,
		Language:   language,
	}
}

func (s *WeatherService) GetWeather(ctx context.Context, location string) (domain.WeatherResponse, error) {
	tracer := otel.Tracer("ms-weatherzip")
	ctx, span := tracer.Start(ctx, "GetWeather")
	span.SetAttributes(attribute.String("location", location))
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	var response domain.WeatherResponse

	encodedLocation := url.QueryEscape(location)
	url := fmt.Sprintf(s.BaseURL, s.ApiKey, encodedLocation, s.Language)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		span.RecordError(err)
		return response, domain.NewFailedToCreateRequestError(err)
	}

	res, err := s.HttpClient.Do(req)
	if err != nil {
		span.RecordError(err)
		return response, domain.NewFailedToMakeRequestError(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusBadRequest {
		var errorResponse struct {
			Error struct {
				Code int `json:"code"`
			} `json:"error"`
		}
		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err == nil {
			if apiErr, exists := weatherErrorCodes[errorResponse.Error.Code]; exists {
				span.SetAttributes(attribute.Int("http.status_code", res.StatusCode))
				span.SetAttributes(attribute.Int("api.error_code", errorResponse.Error.Code))
				span.RecordError(apiErr)
				return response, apiErr
			}
		}
		span.RecordError(err)
		return response, domain.ErrUnexpectedBadRequest
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusBadRequest {
		span.SetAttributes(attribute.Int("http.status_code", res.StatusCode))
		span.RecordError(err)
		return response, domain.NewUnexpectedStatusCodeError(res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		span.RecordError(err)
		return response, domain.NewFailedToDecodeResponseError(err)
	}

	response.Current.TempK = response.Current.TempC + 273

	return response, nil
}
