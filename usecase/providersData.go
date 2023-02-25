package usecase

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/jmoiron/sqlx"
)

type providersDataUsecase struct {
	providersDataRepo domain.ProvidersDataRepo
}

func NewProvidersDataUsecase(pdr domain.ProvidersDataRepo) domain.ProvidersDataUsecase {
	return &providersDataUsecase{
		providersDataRepo: pdr,
	}
}

func (u *providersDataUsecase) Create(pd *domain.ProvidersData, tx *sqlx.Tx) error {
	return u.providersDataRepo.Create(pd, tx)
}

func (u *providersDataUsecase) FindByUsernameAndProvider(pd *domain.ProvidersData, username string, providerID uint) error {
	return u.providersDataRepo.FindByUsernameAndProvider(pd, username, providerID)
}

func (u *providersDataUsecase) GetAllByProvider(userUUID string, providerID uint) ([]domain.ProvidersData, error) {
	return u.providersDataRepo.GetAllByProvider(userUUID, providerID)
}
