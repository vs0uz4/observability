package mock

import (
	"context"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

type MockCepService struct {
	GetLocationFunc func(ctx context.Context, cep string) (domain.CepResponse, error)
}

func (m *MockCepService) GetLocation(ctx context.Context, cep string) (domain.CepResponse, error) {
	return m.GetLocationFunc(ctx, cep)
}
