package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type GithubUser struct {
	BaseModel
	ID        uint      `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	Email     string    `json:"email" db:"email"`
	AvatarURL string    `json:"avatar_url" db:"avatar_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *GithubUser) Create(login string) (*sqlx.DB, error) {
	u.CreatedAt = time.Now()
	db, err := u.createSubDir("github", login)

	if err != nil {
		return nil, err
	}

	_, err = db.NamedExec(`INSERT INTO users (login, email, avatar_url, created_at) 
								VALUES (:login, :email, :avatar_url, :created_at)`, u)

	if err != nil {
		return nil, errors.New("failed to create user " + err.Error())
	}

	// update model
	if err = db.Get(u, "SELECT * FROM users ORDER BY id DESC LIMIT 1"); err != nil {
		return nil, err
	}

	return db, nil
}

func (u *GithubUser) CheckAndUpdateDb(login string) (*sqlx.DB, bool) {
	db, err := u.CheckDb("github", login)
	if err != nil {
		return nil, false
	}

	if err = u.FindByLogin(db, login); err != nil {
		return nil, false
	}

	return db, true
}

func (u *GithubUser) FindByLogin(db *sqlx.DB, login string) error {
	if err := db.Get(u, "SELECT * FROM users WHERE login=$1", login); err != nil {
		return err
	}

	return nil
}

func (u *GithubUser) GetLastToken(db *sqlx.DB) (AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM tokens WHERE user_id=$1 ORDER BY id DESC", u.ID)
	if err != nil {
		return t, err
	}

	return t, err
}

func (u *GithubUser) GetTokenByStr(db *sqlx.DB, token string) (AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM tokens WHERE user_id=$1 AND token=$2 ORDER BY id DESC", u.ID, token)
	if err != nil {
		return t, err
	}

	return t, nil
}
