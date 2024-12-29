package usecase

import (
	"github.com/vs0uz4/observability/ms-searchzip/internal/domain"
	"github.com/vs0uz4/observability/ms-searchzip/internal/service/contract"
)

type locationByCepUsecase struct {
	CepService contract.CepService
}

func NewLocationByCepUsecase(cepService contract.CepService) *locationByCepUsecase {
	return &locationByCepUsecase{
		CepService: cepService,
	}
}

func (uc *locationByCepUsecase) GetLocationByCep(cep string) (domain.CepResponse, error) {
	if len(cep) != 8 || !isNumeric(cep) {
		return domain.CepResponse{}, domain.ErrInvalidZipcode
	}

	location, err := uc.CepService.GetLocation(cep)
	if err != nil {
		return domain.CepResponse{}, err
	}

	return location, nil
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
