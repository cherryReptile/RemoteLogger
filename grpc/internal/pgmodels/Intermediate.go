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
	_, err := db.NamedExec(`INSERT INTO intermediate (user_id, provider_id, created_at) 
								VALUES (:user_id, :provider_id, :created_at)`, i)

	if err != nil {
		return err
	}

	// update model
	if err = db.Get(i, "SELECT * FROM intermediate ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}
