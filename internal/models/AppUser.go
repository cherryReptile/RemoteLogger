package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type AppUser struct {
	BaseModel
	ID        uint      `json:"ID" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *AppUser) Create(email string) (*sqlx.DB, error) {
	u.CreatedAt = time.Now()
	db, err := u.createSubDir("app", email)

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

func (u *AppUser) CheckAndUpdateDb(email string) (*sqlx.DB, bool) {
	db, err := u.CheckDb("app", email)
	if err != nil {
		return db, false
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

func (u *AppUser) GetLastToken(db *sqlx.DB) (AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM tokens WHERE user_id=$1 ORDER BY id DESC", u.ID)
	if err != nil {
		return t, err
	}

	return t, err
}

func (u *AppUser) GetTokenByStr(db *sqlx.DB, token string) (AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM tokens WHERE user_id=$1 AND token=$2 ORDER BY id DESC", u.ID, token)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (u *AppUser) GetUniqueRaw() string {
	return u.Email
}
