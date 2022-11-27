package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type AccessToken struct {
	ID        uint      `json:"id" db:"id""`
	Token     string    `json:"access_token" db:"token"`
	UserID    uint      `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (t *AccessToken) Create(db *sqlx.DB) error {
	t.CreatedAt = time.Now()

	_, err := db.NamedExec(`INSERT INTO tokens (token, user_id, created_at) 
								VALUES (:token, :user_id, :created_at)`, t)

	if err != nil {
		return errors.New("failed to create token " + err.Error())
	}

	// update model
	if err = db.Get(t, "SELECT * FROM tokens ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}

func (t *AccessToken) GetByToken(db *sqlx.DB, token string) error {
	if err := db.Get(t, "SELECT * FROM tokens WHERE token=$1 ORDER BY id DESC LIMIT 1", token); err != nil {
		return err
	}

	return nil
}

func (t *AccessToken) Delete(db *sqlx.DB) error {
	_, err := db.NamedExec("DELETE FROM tokens WHERE id=:id", t)
	return err
}
