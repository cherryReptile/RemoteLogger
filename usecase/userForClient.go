package usecase

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/jmoiron/sqlx"
)

type clientUserUsecase struct {
	clientUserRepo domain.ClientUserRepo
}

func NewUserAndProfileUsecase(userAndProfileRepo domain.ClientUserRepo) domain.ClientUserUsecase {
	return &clientUserUsecase{
		clientUserRepo: userAndProfileRepo,
	}
}

func (u *clientUserUsecase) GetUserWithProfile(clientUser *domain.ClientUser, userID string) error {
	return u.clientUserRepo.GetUserWithProfile(clientUser, userID)
}

func (u *clientUserUsecase) GetAuthClientUser(clientUser *domain.ClientUser, userID, token string) error {
	return u.clientUserRepo.GetAuthClientUser(clientUser, userID, token)
}

func (u *clientUserUsecase) GetAllWithOrderBy(field, orderBy string) (*sqlx.Rows, error) {
	return u.clientUserRepo.GetAllWithOrderBy(field, orderBy)
}

func (u *clientUserUsecase) GetAllWithOrderByAndFilter(filter map[string]string, field, orderBy string) (*sqlx.Rows, error) {
	return u.clientUserRepo.GetAllWithOrderByAndFilter(filter, field, orderBy)
}
