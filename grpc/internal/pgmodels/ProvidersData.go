package pgmodels

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"time"
)

type ProvidersData struct {
	ID         uint            `json:"id" db:"id"`
	UserData   json.RawMessage `json:"user_data" db:"user_data"`
	UserID     string          `json:"user_id" db:"user_id"`
	ProviderID uint            `json:"provider_id" db:"provider_id"`
	Username   string          `json:"username" db:"username"`
	CreatedAt  time.Time       `json:"create_at" db:"created_at"`
}

func (p *ProvidersData) Create(db *sqlx.DB) error {
	p.CreatedAt = time.Now()

	if len(p.UserData) > 0 {
		json, err := helpers.TrimJson(p.UserData)
		if err != nil {
			return err
		}
		p.UserData = json
	}
	_, err := db.NamedExec(`INSERT INTO users_providers_data (user_data, user_id, provider_id, username, created_at) 
								VALUES (:user_data, :user_id, :provider_id, :username, :created_at)`, p)

	if err != nil {
		return err
	}

	// update model
	if err = db.Get(p, "SELECT * FROM users_providers_data ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}

func (p *ProvidersData) FindByUsernameAndProvider(db *sqlx.DB, username string, providerID uint) error {
	if err := db.Get(p, "SELECT * FROM users_providers_data WHERE username=$1 AND provider_id=$2 LIMIT 1", username, providerID); err != nil {
		return err
	}

	return nil
}

func (p *ProvidersData) FindByUserUUIDAndProviderID(db *sqlx.DB, userUUID string, providerID uint) error {
	if err := db.Get(p, "SELECT * FROM users_providers_data WHERE user_id=$1 AND provider_id=$2", userUUID, providerID); err != nil {
		return err
	}
	return nil
}
