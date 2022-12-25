package pgmodels

import (
	"github.com/jmoiron/sqlx"
)

type AuthProvider struct {
	BaseModel
	ID        uint   `db:"id"`
	Provider  string `db:"provider"`
	UniqueKey string `db:"unique_key"`
}

func (a *AuthProvider) GetByProvider(db *sqlx.DB, provider string) error {
	if err := db.Get(a, "SELECT * FROM auth_providers WHERE provider=$1 LIMIT 1", provider); err != nil {
		return err
	}

	return nil
}
