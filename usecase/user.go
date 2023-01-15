package usecase

import (
	"github.com/cherryReptile/WS-AUTH/domain"
)

type userUsecase struct {
	userRepo domain.UserRepo
}

func NewUserUsecase(ur domain.UserRepo) domain.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (u *userUsecase) Create(user *domain.User) error {
	return u.userRepo.Create(user)
}

func (u *userUsecase) Find(user *domain.User, uuid string) error {
	return u.userRepo.Find(user, uuid)
}

func (u *userUsecase) FindByLoginAndProvider(user *domain.User, username, provider string) error {
	return u.userRepo.FindByLoginAndProvider(user, username, provider)
}

func (u *userUsecase) GetTokenByStr(user *domain.User, tokenStr string) (*domain.AuthToken, error) {
	return u.userRepo.GetTokenByStr(user, tokenStr)
}
