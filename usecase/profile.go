package usecase

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/jmoiron/sqlx"
)

type profileUsecase struct {
	profileRepo domain.ProfileRepo
}

func NewProfileUsecase(pr domain.ProfileRepo) domain.ProfileUsecase {
	return &profileUsecase{
		profileRepo: pr,
	}
}

func (u *profileUsecase) Create(profile *domain.Profile, tx *sqlx.Tx) error {
	return u.profileRepo.Create(profile, tx)
}

func (u *profileUsecase) FindByUserUUID(profile *domain.Profile, userUUID string) error {
	return u.profileRepo.FindByUserUUID(profile, userUUID)
}

func (u *profileUsecase) Update(profile *domain.Profile) error {
	return u.profileRepo.Update(profile)
}

func (u *profileUsecase) Delete(profile *domain.Profile) error {
	return u.profileRepo.Delete(profile)
}
