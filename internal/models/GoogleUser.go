package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/internal/sqlite"
	"os"
	"time"
)

type GoogleUser struct {
	ID        uint      `json:"user_id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Picture   string    `json:"picture" db:"picture"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *GoogleUser) Create(email string) (*sqlx.DB, error) {
	u.CreatedAt = time.Now()
	db, err := u.createSubDir(email)

	if err != nil {
		return nil, err
	}

	_, err = db.NamedExec(`INSERT INTO users (email, picture, created_at) 
								VALUES (:email, :picture, :created_at)`, u)

	if err != nil {
		return nil, errors.New("failed to create user " + err.Error())
	}

	// update model
	if err = db.Get(u, "SELECT * FROM users ORDER BY id DESC LIMIT 1"); err != nil {
		return nil, err
	}

	return db, nil
}

func (u *GoogleUser) CheckDb(login string) (*sqlx.DB, bool) {
	path := "./storage/users/google/" + login
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return nil, false
	}

	db, err := sqlite.GetDb("google", login)
	if err != nil {
		return nil, false
	}

	if err = u.FindByEmail(db, login); err != nil {
		return nil, false
	}

	return db, true
}

func (u *GoogleUser) FindByEmail(db *sqlx.DB, email string) error {
	if err := db.Get(u, "SELECT * FROM users WHERE email=$1", email); err != nil {
		return err
	}

	return nil
}

func (u *GoogleUser) GetTokenByStr(db *sqlx.DB, token string) (AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM tokens WHERE user_id=$1 AND token=$2 ORDER BY id DESC", u.ID, token)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (u *GoogleUser) createSubDir(email string) (*sqlx.DB, error) {
	path := "./storage/users/google/" + email
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, err
	}

	db, err := sqlite.GetDb("google", email)
	if err != nil {
		return nil, err
	}

	err = sqlite.SetDefaultSchema(db, "google")
	if err != nil {
		return nil, err
	}

	return db, nil
}
