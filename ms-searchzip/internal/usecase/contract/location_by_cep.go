package contract

import "github.com/vs0uz4/observability/ms-searchzip/internal/domain"

type LocationByCepUsecase interface {
	GetLocationByCep(cep string) (domain.CepResponse, error)
}
