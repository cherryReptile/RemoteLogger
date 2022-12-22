package pgmodels

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type AuthProvider struct {
	BaseModel
	ID        uint   `db:"id"`
	Provider  string `db:"provider"`
	UniqueKey string `db:"unique_key"`
}

func (a *AuthProvider) SetAuthProviders(db *sqlx.DB, provider string) error {
	switch provider {
	case "app":
		a.UniqueKey = "email"
	case "github":
		a.UniqueKey = "login"
	case "google":
		a.UniqueKey = "email"
	case "telegram":
		a.UniqueKey = "username"
	default:
		return errors.New("unknown auth provider")
	}
	a.Provider = provider

	_, err := db.NamedExec(`INSERT INTO auth_providers (provider, unique_key) 
								VALUES (:provider, :unique_key)`, a)

	if err != nil {
		return err
	}

	// update model
	if err = db.Get(a, "SELECT * FROM auth_providers ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}

func (a *AuthProvider) GetByProvider(db *sqlx.DB, provider string) error {
	if err := db.Get(a, "SELECT * FROM auth_providers WHERE provider=$1", provider); err != nil {
		return err
	}

	return nil
}
