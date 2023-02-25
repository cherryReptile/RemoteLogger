package usecase

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/jmoiron/sqlx"
)

type usersProvidersUsecase struct {
	usersProvidersRepo domain.UsersProvidersRepo
}

func NewUsersProvidersUsecase(upr domain.UsersProvidersRepo) domain.UsersProvidersUsecase {
	return &usersProvidersUsecase{
		usersProvidersRepo: upr,
	}
}

func (u *usersProvidersUsecase) Create(up *domain.UsersProviders, userUUID string, providerID uint, tx *sqlx.Tx) error {
	return u.usersProvidersRepo.Create(up, userUUID, providerID, tx)
}
