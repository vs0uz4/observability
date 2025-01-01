package contract

import (
	"context"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

type CepService interface {
	GetLocation(ctx context.Context, cep string) (domain.CepResponse, error)
}
