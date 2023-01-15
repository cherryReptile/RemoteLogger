package repository

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/cherryReptile/WS-AUTH/internal/helpers"
	"github.com/jmoiron/sqlx"
	"time"
)

type profileRepository struct {
	db *sqlx.DB
}

func NewProfileRepository(db *sqlx.DB) domain.ProfileRepo {
	return &profileRepository{
		db: db,
	}
}

func (r *profileRepository) Create(profile *domain.Profile) error {
	profile.CreatedAt = time.Now()

	if len(profile.OtherData) > 0 {
		json, err := helpers.TrimJson(profile.OtherData)
		if err != nil {
			return err
		}
		profile.OtherData = json
	}

	_, err := r.db.NamedExec(`INSERT INTO user_profiles (first_name, last_name, other_data, address, user_id, created_at) 
								VALUES (:first_name, :last_name, :other_data, :address, :user_id, :created_at)`, profile)

	if err != nil {
		return err
	}

	// update model
	if err = r.db.Get(profile, "SELECT * FROM user_profiles ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}

func (r *profileRepository) FindByUserUUID(profile *domain.Profile, userUUID string) error {
	if err := r.db.Get(profile, "SELECT * FROM user_profiles WHERE user_id=$1", userUUID); err != nil {
		return err
	}
	return nil
}

func (r *profileRepository) Update(profile *domain.Profile) error {
	_, err := r.db.Exec("UPDATE user_profiles SET first_name=$1, last_name=$2, other_data=$3, address=$4", profile.FirstName, profile.LastName, profile.OtherData, profile.Address)
	if err != nil {
		return err
	}

	return nil
}

func (r *profileRepository) Delete(profile *domain.Profile) error {
	if _, err := r.db.Exec("DELETE FROM user_profiles WHERE id=$1", profile.ID); err != nil {
		return err
	}

	return nil
}
