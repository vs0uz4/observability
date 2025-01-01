package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/vs0uz4/observability/ms-inputvalidate/internal/domain"
	"github.com/vs0uz4/observability/ms-inputvalidate/internal/service/contract"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
)

var _ contract.WeatherService = (*WeatherService)(nil)

var weatherErrorMessages = map[string]struct {
	Error error
}{
	"invalid zipcode":       {Error: domain.ErrInvalidZipcode},
	"can not find zipcode":  {Error: domain.ErrZipcodeNotFound},
	"internal server error": {Error: domain.ErrInternalServer},
}

type WeatherService struct {
	HttpClient contract.HttpClient
	BaseURL    string
}

func NewWeatherService(client *http.Client, baseURL string) *WeatherService {
	return &WeatherService{
		HttpClient: client,
		BaseURL:    baseURL,
	}
}

func (s *WeatherService) GetWeather(ctx context.Context, cep string) (domain.WeatherResponse, error) {
	tracer := otel.Tracer("ms-inputvalidate")
	ctx, span := tracer.Start(ctx, "GetWeather")
	span.SetAttributes(attribute.String("cep", cep))
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	var response domain.WeatherResponse

	url := fmt.Sprintf(s.BaseURL, cep)
	span.SetAttributes(
		attribute.String("peer.service", "ms-weatherzip"),
		attribute.String("http.url", url),
		attribute.String("http.method", http.MethodGet),
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		span.RecordError(err)
		return response, domain.NewFailedToCreateRequestError(err)
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	res, err := s.HttpClient.Do(req)
	if err != nil {
		span.RecordError(err)
		return response, domain.NewFailedToMakeRequestError(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err := s.parseErrorResponse(res)

		span.SetAttributes(attribute.Int("http.status_code", res.StatusCode))
		span.RecordError(err)
		return response, err
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		span.RecordError(err)
		return response, domain.NewFailedToDecodeResponseError(err)
	}

	return response, nil
}

func (s *WeatherService) parseErrorResponse(res *http.Response) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return domain.ErrReadResponseBodyData
	}

	message := strings.TrimSpace(string(body))

	var exists bool

	if apiErr, exists := weatherErrorMessages[message]; exists {
		return apiErr.Error
	}

	if res.StatusCode == http.StatusUnprocessableEntity && !exists {
		return domain.ErrUnexpectedUnprocessableEntity
	}

	if res.StatusCode == http.StatusNotFound && !exists {
		return domain.ErrUnexpectedNotFound
	}

	if res.StatusCode == http.StatusInternalServerError && !exists {
		return domain.ErrUnexpectedInternalServer
	}

	return domain.NewUnexpectedWeatherServiceError(
		domain.ErrWeatherService,
		res.StatusCode,
		message,
	)
}
