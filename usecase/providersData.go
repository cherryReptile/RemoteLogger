package usecase

import "github.com/cherryReptile/WS-AUTH/domain"

type providersDataUsecase struct {
	providersDataRepo domain.ProvidersDataRepo
}

func NewProvidersDataUsecase(pdr domain.ProvidersDataRepo) domain.ProvidersDataUsecase {
	return &providersDataUsecase{
		providersDataRepo: pdr,
	}
}

func (u *providersDataUsecase) Create(pd *domain.ProvidersData) error {
	return u.providersDataRepo.Create(pd)
}

func (u *providersDataUsecase) FindByUsernameAndProvider(pd *domain.ProvidersData, username string, providerID uint) error {
	return u.providersDataRepo.FindByUsernameAndProvider(pd, username, providerID)
}

func (u *providersDataUsecase) GetAllByProvider(userUUID string, providerID uint) ([]domain.ProvidersData, error) {
	return u.providersDataRepo.GetAllByProvider(userUUID, providerID)
}
