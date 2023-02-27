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

func (r *profileRepository) Create(profile *domain.Profile, tx *sqlx.Tx) error {
	var err error
	create := `INSERT INTO user_profiles (first_name, last_name, other_data, address, user_id, created_at) 
								VALUES (:first_name, :last_name, :other_data, :address, :user_id, :created_at)`
	get := "SELECT * FROM user_profiles ORDER BY id DESC LIMIT 1"
	profile.CreatedAt = time.Now()

	if len(profile.OtherData) > 0 {
		json, err := helpers.TrimJson(profile.OtherData)
		if err != nil {
			return err
		}
		profile.OtherData = json
	}

	if tx != nil {
		_, err = tx.NamedExec(create, profile)

		if err != nil {
			return Rollback(err, tx)
		}

		if err = tx.Get(profile, get); err != nil {
			return Rollback(err, tx)
		}

		return nil
	}

	_, err = r.db.NamedExec(create, profile)

	if err != nil {
		return err
	}

	return r.db.Get(profile, get)
}

func (r *profileRepository) FindByUserUUID(profile *domain.Profile, userUUID string) error {
	if err := r.db.Get(profile, "SELECT * FROM user_profiles WHERE user_id=$1", userUUID); err != nil {
		return err
	}
	return nil
}

func (r *profileRepository) Update(profile *domain.Profile) error {
	_, err := r.db.Exec("UPDATE user_profiles SET first_name=$1, last_name=$2, other_data=$3, address=$4 WHERE user_id=$5", profile.FirstName, profile.LastName, profile.OtherData, profile.Address, profile.UserID)
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
