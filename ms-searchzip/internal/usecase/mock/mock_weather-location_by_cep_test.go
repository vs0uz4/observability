package mock

import (
	"errors"
	"testing"

	"github.com/vs0uz4/observability/ms-searchzip/internal/domain"
)

func TestMockLocationByCepUsecase(t *testing.T) {
	mockUsecase := &MockLocationByCepUsecase{
		GetLocationByCepFunc: func(cep string) (domain.CepResponse, error) {
			if cep == "12345678" {
				return domain.CepResponse{
					Cep:        "12345678",
					Logradouro: "Rua Teste",
					Bairro:     "Bairro Teste",
					Localidade: "Cidade Teste",
					Uf:         "UF",
				}, nil
			}
			return domain.CepResponse{}, errors.New("invalid cep")
		},
	}

	t.Run("Success", func(t *testing.T) {
		resp, err := mockUsecase.GetLocationByCep("12345678")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if resp.Cep != "12345678" {
			t.Errorf("Expected Cep '12345678', got %v", resp.Cep)
		}

		if resp.Logradouro != "Rua Teste" {
			t.Errorf("Expected Logradouro 'Rua Teste', got %v", resp.Logradouro)
		}

		if resp.Bairro != "Bairro Teste" {
			t.Errorf("Expected Bairro 'Bairro Teste', got %v", resp.Bairro)
		}

		if resp.Localidade != "Cidade Teste" {
			t.Errorf("Expected Localidade 'Cidade Teste', got %v", resp.Localidade)
		}

		if resp.Uf != "UF" {
			t.Errorf("Expected Uf 'UF', got %v", resp.Uf)
		}
	})

	t.Run("Failure", func(t *testing.T) {
		_, err := mockUsecase.GetLocationByCep("00000000")
		if err == nil || err.Error() != "invalid cep" {
			t.Errorf("Expected error 'invalid cep', got %v", err)
		}
	})
}
