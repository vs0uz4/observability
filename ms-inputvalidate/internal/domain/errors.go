package domain

import (
	"errors"
	"fmt"
)

var (
	ErrCepService                    = errors.New("cep service error")
	ErrWeatherService                = errors.New("unexpected response from weather service")
	ErrZipcodeNotFound               = errors.New("zipcode not found")
	ErrInvalidZipcode                = errors.New("invalid zipcode")
	ErrLocationNotFound              = errors.New("location not found")
	ErrUnexpectedUrl                 = errors.New("unexpected url")
	ErrUnexpectedNotFound            = errors.New("unexpected not found error")
	ErrUnexpectedBadRequest          = errors.New("unexpected bad request error")
	ErrUnexpectedInternalServer      = errors.New("unexpected internal server error")
	ErrUnexpectedUnprocessableEntity = errors.New("unexpected unprocessable entity error")
	ErrParameterNotProvided          = errors.New("parameter 'q' not provided")
	ErrApiUrlIsInvalid               = errors.New("api request url is invalid")
	ErrJsonBodyIsInvalid             = errors.New("invalid json body in bulk request")
	ErrTooManyLocations              = errors.New("too many locations in bulk request")
	ErrInternalApplication           = errors.New("internal application error")
	ErrInternalServer                = errors.New("internal server error")
	ErrInvalidLocationData           = errors.New("invalid location data")
	ErrInvalidCurrentData            = errors.New("invalid current weather data")
	ErrInvalidZipCodeData            = errors.New("invalid zipcode data")
	ErrInvalidStreetData             = errors.New("invalid street data")
	ErrInvalidNeighborhoodData       = errors.New("invalid neighborhood data")
	ErrInvalidFederativeUnitData     = errors.New("invalid federative unit data")
	ErrReadResponseBodyData          = errors.New("failed to read response body")
)

func NewUnexpectedStatusCodeError(statusCode int) error {
	return fmt.Errorf("unexpected status code: %d", statusCode)
}

func NewUnexpectedWeatherServiceError(err error, statusCode int, responseBody string) error {
	return fmt.Errorf("unexpected error (HTTP %d): %s - details %w", statusCode, responseBody, err)
}

func NewFailedToCreateRequestError(err error) error {
	return fmt.Errorf("failed to create request: %w", err)
}

func NewFailedToMakeRequestError(err error) error {
	return fmt.Errorf("failed to make request: %w", err)
}

func NewFailedToDecodeResponseError(err error) error {
	return fmt.Errorf("failed to decode response: %w", err)
}
