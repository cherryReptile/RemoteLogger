package models

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Profile struct {
	ID        uint           `json:"id" db:"id"`
	FirstName string         `json:"first_name" db:"first_name"`
	LastName  string         `json:"last_name" db:"last_name"`
	OtherData sql.NullString `json:"other_data" db:"other_data"`
	UserID    string         `json:"user_id" db:"user_id"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
}

func (p *Profile) Create(db *sqlx.DB) error {
	p.CreatedAt = time.Now()
	_, err := db.NamedExec(`INSERT INTO user_profiles (first_name, last_name, other_data, user_id, created_at) 
								VALUES (:first_name, :last_name, :other_data, :user_id, :created_at)`, p)

	if err != nil {
		return err
	}

	// update model
	if err = db.Get(p, "SELECT * FROM user_profiles ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}
