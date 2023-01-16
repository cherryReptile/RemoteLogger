package repository

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/cherryReptile/WS-AUTH/internal/helpers"
	"github.com/jmoiron/sqlx"
	"time"
)

type providersDataRepo struct {
	db *sqlx.DB
}

func NewProvidersDataRepo(db *sqlx.DB) domain.ProvidersDataRepo {
	return &providersDataRepo{
		db: db,
	}
}

func (r *providersDataRepo) Create(pd *domain.ProvidersData) error {
	pd.CreatedAt = time.Now()

	if len(pd.UserData) > 0 {
		json, err := helpers.TrimJson(pd.UserData)
		if err != nil {
			return err
		}
		pd.UserData = json
	}
	_, err := r.db.NamedExec(`INSERT INTO users_providers_data (user_data, user_id, provider_id, username, created_at) 
								VALUES (:user_data, :user_id, :provider_id, :username, :created_at)`, pd)

	if err != nil {
		return err
	}

	// update model
	if err = r.db.Get(pd, "SELECT * FROM users_providers_data ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}

func (r *providersDataRepo) FindByUsernameAndProvider(pd *domain.ProvidersData, username string, providerID uint) error {
	if err := r.db.Get(pd, "SELECT * FROM users_providers_data WHERE username=$1 AND provider_id=$2 LIMIT 1", username, providerID); err != nil {
		return err
	}

	return nil
}

func (r *providersDataRepo) GetAllByProvider(userUUID string, providerID uint) ([]domain.ProvidersData, error) {
	var ps []domain.ProvidersData
	if err := r.db.Select(&ps, "SELECT * FROM users_providers_data WHERE user_id=$1 AND provider_id=$2", userUUID, providerID); err != nil {
		return nil, err
	}

	return ps, nil
}
