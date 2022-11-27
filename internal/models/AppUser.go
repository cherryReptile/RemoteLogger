package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/internal/sqlite"
	"os"
	"time"
)

const dbAppPath = "./storage/users/app/"

type AppUser struct {
	ID        uint      `json:"ID" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *AppUser) Create(email string) (*sqlx.DB, error) {
	u.CreatedAt = time.Now()
	db, err := u.createSubDir(email)

	if err != nil {
		return nil, err
	}

	_, err = db.NamedExec(`INSERT INTO users (email, password, created_at) 
								VALUES (:email, :password, :created_at)`, u)

	if err != nil {
		return nil, errors.New("failed to create user " + err.Error())
	}

	// update model
	if err = db.Get(u, "SELECT * FROM users ORDER BY id DESC LIMIT 1"); err != nil {
		return nil, err
	}

	return db, nil
}

func (u *AppUser) CheckDb(email string) (*sqlx.DB, bool) {
	path := dbAppPath + email
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return nil, false
	}

	db, err := sqlite.GetDb("app", email)
	if err != nil {
		return nil, false
	}

	if err = u.FindByEmail(db, email); err != nil {
		return nil, false
	}

	return db, true
}

func (u *AppUser) FindByEmail(db *sqlx.DB, email string) error {
	if err := db.Get(u, "SELECT * FROM users WHERE email=$1", email); err != nil {
		return err
	}

	return nil
}

func (u *AppUser) GetAccessToken(db *sqlx.DB) (AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM tokens WHERE user_id=$1 ORDER BY id DESC", u.ID)
	if err != nil {
		return t, err
	}

	return t, err
}

func (u *AppUser) createSubDir(email string) (*sqlx.DB, error) {
	path := dbAppPath + email
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, err
	}

	db, err := sqlite.GetDb("app", email)
	if err != nil {
		return nil, err
	}

	err = sqlite.SetDefaultSchema(db, "app")
	if err != nil {
		return nil, err
	}

	return db, nil
}
