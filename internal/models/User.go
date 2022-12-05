package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	BaseModel
	ID           uint      `json:"id"`
	UniqueRaw    string    `json:"unique_raw" db:"unique_raw"`
	Password     string    `json:"password" db:"password"`
	AuthorizedBy string    `json:"authorized_by" db:"authorized_by"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

func (u *User) Create(db *sqlx.DB) error {
	u.CreatedAt = time.Now()

	_, err := db.NamedExec(`INSERT INTO users (unique_raw, password, authorized_by, created_at) 
								VALUES (:unique_raw, :password, :authorized_by, :created_at)`, u)

	if err != nil {
		//return errors.New("failed to create user " + err.Error())
		return err
	}

	// update model
	if err = db.Get(u, "SELECT * FROM users ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}

func (u *User) FindByUniqueAndService(db *sqlx.DB, unique, service string) error {
	if err := db.Get(u, "SELECT * FROM users WHERE unique_raw=$1 AND authorized_by=$2", unique, service); err != nil {
		return err
	}

	return nil
}

func (u *User) GetUniqueRaw() string {
	return u.UniqueRaw
}

func (u *User) GetTokenByStr(db *sqlx.DB, token string) (*AccessToken, error) {
	t := AccessToken{}
	err := db.Get(&t, "SELECT * FROM access_tokens WHERE user_id=$1 AND token=$2 ORDER BY id DESC", u.ID, token)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
