package pgmodels

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Intermediate struct {
	BaseModel
	ID         uint      `json:"id" db:"id"`
	UserID     string    `json:"user_id" db:"user_id"`
	ProviderID uint      `json:"provider_id" db:"provider_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

func (i *Intermediate) Create(db *sqlx.DB, uuid string, providerID uint) error {
	i.UserID = uuid
	i.ProviderID = providerID
	i.CreatedAt = time.Now()
	_, err := db.NamedExec(`INSERT INTO users_providers (user_id, provider_id, created_at) 
								VALUES (:user_id, :provider_id, :created_at)`, i)

	if err != nil {
		return err
	}

	// update model
	if err = db.Get(i, "SELECT * FROM users_providers ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}

func (i *Intermediate) Find(db *sqlx.DB, userUUID string, providerID uint) error {
	if err := db.Get(i, "SELECT * FROM users_providers WHERE user_id=$1 AND provider_id=$2 LIMIT 1", userUUID, providerID); err != nil {
		return err
	}
	return nil
}
