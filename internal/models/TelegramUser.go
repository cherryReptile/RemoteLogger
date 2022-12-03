package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

type TelegramUser struct {
	BaseModel
	ID        uint      `json:"app_id"`
	TgID      uint      `json:"id" db:"tg_id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Username  string    `json:"username" db:"username"`
	PhotoURL  string    `json:"photo_url" db:"photo_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *TelegramUser) Create(email string) (*sqlx.DB, error) {
	u.CreatedAt = time.Now()
	db, err := u.createSubDir("telegram", email)

	if err != nil {
		return nil, err
	}

	_, err = db.NamedExec(`INSERT INTO users (tg_id, first_name, last_name, username, photo_url, created_at) 
								VALUES (:tg_id, :first_name, :last_name, :username, :photo_url, :created_at)`, u)

	if err != nil {
		return nil, errors.New("failed to create user " + err.Error())
	}

	// update model
	if err = db.Get(u, "SELECT * FROM users ORDER BY id DESC LIMIT 1"); err != nil {
		return nil, err
	}

	return db, nil
}

func (u *TelegramUser) CheckAndUpdateDb(login string) (*sqlx.DB, bool) {
	db, err := u.CheckDb("telegram", login)
	if err != nil {
		return nil, false
	}

	if err = u.FindByUsername(db, login); err != nil {
		return nil, false
	}

	return db, true
}

func (u *TelegramUser) FindByUsername(db *sqlx.DB, username string) error {
	return nil
}

func (u *TelegramUser) GetTokenByStr(db *sqlx.DB, token string) (AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM tokens WHERE user_id=$1 AND token=$2 ORDER BY id DESC", u.ID, token)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (u *TelegramUser) GetUniqueRaw() string {
	return u.Username
}
