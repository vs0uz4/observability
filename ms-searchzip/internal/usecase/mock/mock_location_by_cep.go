package mock

import (
	"github.com/vs0uz4/observability/ms-searchzip/internal/domain"
)

type MockLocationByCepUsecase struct {
	GetLocationByCepFunc func(cep string) (domain.CepResponse, error)
}

func (m *MockLocationByCepUsecase) GetLocationByCep(cep string) (domain.CepResponse, error) {
	return m.GetLocationByCepFunc(cep)
}
