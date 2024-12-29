package contract

import "github.com/vs0uz4/observability/ms-searchzip/internal/domain"

type CepService interface {
	GetLocation(cep string) (domain.CepResponse, error)
}
