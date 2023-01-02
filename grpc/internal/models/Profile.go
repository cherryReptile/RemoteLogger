package models

import (
	"database/sql"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/internal/helpers"
	"time"
)

type Profile struct {
	ID        uint            `json:"id" db:"id"`
	FirstName sql.NullString  `json:"first_name" db:"first_name"`
	LastName  sql.NullString  `json:"last_name" db:"last_name"`
	OtherData json.RawMessage `json:"other_data" db:"other_data"`
	Address   sql.NullString  `json:"address" db:"address"`
	UserID    string          `json:"user_id" db:"user_id"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
}

func (p *Profile) Create(db *sqlx.DB) error {
	p.CreatedAt = time.Now()

	if len(p.OtherData) > 0 {
		json, err := helpers.TrimJson(p.OtherData)
		if err != nil {
			return err
		}
		p.OtherData = json
	}

	_, err := db.NamedExec(`INSERT INTO user_profiles (first_name, last_name, other_data, address, user_id, created_at) 
								VALUES (:first_name, :last_name, :other_data, :address, :user_id, :created_at)`, p)

	if err != nil {
		return err
	}

	// update model
	if err = db.Get(p, "SELECT * FROM user_profiles ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}

func (p *Profile) FindByUserUUID(db *sqlx.DB, userUUID string) error {
	if err := db.Get(p, "SELECT * FROM user_profiles WHERE user_id=$1", userUUID); err != nil {
		return err
	}
	return nil
}

func (p *Profile) Update(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE user_profiles SET first_name=$1, last_name=$2, other_data=$3, address=$4", p.FirstName, p.LastName, p.OtherData, p.Address)
	if err != nil {
		return err
	}

	return nil
}

func (p *Profile) Delete(db *sqlx.DB) error {
	if _, err := db.Exec("DELETE FROM user_profiles WHERE id=$1", p.ID); err != nil {
		return err
	}

	return nil
}
