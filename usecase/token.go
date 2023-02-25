package usecase

import (
	"github.com/cherryReptile/WS-AUTH/domain"
)

type tokenUsecase struct {
	tokeRepo domain.AuthTokenRepo
}

func NewTokenUsecase(tp domain.AuthTokenRepo) domain.AuthTokenUsecase {
	return &tokenUsecase{
		tokeRepo: tp,
	}
}

func (u *tokenUsecase) Create(t *domain.AuthToken) error {
	return u.tokeRepo.Create(t)
}

func (u *tokenUsecase) GetByToken(token *domain.AuthToken, tokenStr string) error {
	return u.tokeRepo.GetByToken(token, tokenStr)
}

func (u *tokenUsecase) Delete(t *domain.AuthToken) error {
	return u.tokeRepo.Delete(t)
}
