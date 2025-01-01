package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
	"github.com/vs0uz4/observability/ms-weatherzip/internal/service/contract"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var _ contract.CepService = (*CepService)(nil)

type CepService struct {
	HttpClient contract.HttpClient
	BaseURL    string
}

func NewCepService(client *http.Client, baseURL string) *CepService {
	return &CepService{
		HttpClient: client,
		BaseURL:    baseURL,
	}
}

func (s *CepService) GetLocation(ctx context.Context, cep string) (domain.CepResponse, error) {
	tracer := otel.Tracer("ms-weatherzip")
	ctx, span := tracer.Start(ctx, "GetLocation")
	span.SetAttributes(attribute.String("cep", cep))
	defer span.End()

	ctx, cancel := context.WithTimeout(ctx, 5000*time.Millisecond)
	defer cancel()

	var response domain.CepResponse
	var raw map[string]interface{}

	url := fmt.Sprintf(s.BaseURL, cep)
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
		span.RecordError(err)
		return response, domain.ErrInvalidZipcode
	}

	if res.StatusCode != http.StatusOK {
		span.SetAttributes(attribute.Int("http.status_code", res.StatusCode))
		span.RecordError(err)
		return response, domain.NewUnexpectedStatusCodeError(res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
		span.RecordError(err)
		return response, domain.NewFailedToDecodeResponseError(err)
	}

	if err := response.PopulateFromMap(raw); err != nil {
		if err.Error() == "zipcode not found" {
			span.RecordError(err)
			return response, domain.ErrZipcodeNotFound
		}
		span.RecordError(err)
		return response, domain.NewFailedToMapResponseError(err)
	}

	return response, nil
}
