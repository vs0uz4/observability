package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/vs0uz4/observability/ms-searchzip/internal/domain"
	"github.com/vs0uz4/observability/ms-searchzip/internal/service/contract"
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

func (s *WeatherService) GetWeather(cep string) (domain.WeatherResponse, error) {
	var response domain.WeatherResponse

	url := fmt.Sprintf(s.BaseURL, cep)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, domain.NewFailedToCreateRequestError(err)
	}

	res, err := s.HttpClient.Do(req)
	if err != nil {
		return response, domain.NewFailedToMakeRequestError(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response, s.parseErrorResponse(res)
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
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
