package domain

import (
	"database/sql"
	"encoding/json"
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
