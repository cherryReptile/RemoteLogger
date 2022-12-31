package models

import (
	"github.com/jmoiron/sqlx"
)

type Provider struct {
	BaseModel
	ID        uint   `db:"id"`
	Provider  string `db:"provider"`
	UniqueKey string `db:"unique_key"`
}

func (a *Provider) GetByProvider(db *sqlx.DB, provider string) error {
	if err := db.Get(a, "SELECT * FROM providers WHERE provider=$1 LIMIT 1", provider); err != nil {
		return err
	}

	return nil
}
