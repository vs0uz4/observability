package mock

import (
	"context"
	"testing"

	"github.com/vs0uz4/observability/ms-weatherzip/internal/domain"
)

func TestMockCepService(t *testing.T) {
	mock := MockCepService{
		GetLocationFunc: func(ctx context.Context, cep string) (domain.CepResponse, error) {
			if cep == "12345678" {
				return domain.CepResponse{Cep: "12345678"}, nil
			}
			return domain.CepResponse{}, domain.ErrZipcodeNotFound
		},
	}

	ctx := context.Background()

	t.Run("Zipcode is Valid", func(t *testing.T) {
		response, err := mock.GetLocation(ctx, "12345678")
		if response.Cep != "12345678" || err != nil {
			t.Errorf("Expected Cep: 12345678, got: %v, err: %v", response.Cep, err)
		}
	})

	t.Run("Zipcode Not Found", func(t *testing.T) {
		_, err := mock.GetLocation(ctx, "00000000")
		if err != domain.ErrZipcodeNotFound {
			t.Errorf("Expected error: %v, got: %v", domain.ErrZipcodeNotFound, err)
		}
	})
}
