package repository

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/jmoiron/sqlx"
)

type providerRepository struct {
	db *sqlx.DB
}

func NewProviderRepository(db *sqlx.DB) domain.ProviderRepo {
	return &providerRepository{
		db: db,
	}
}

func (r *providerRepository) GetByProvider(provider *domain.Provider, name string) error {
	if err := r.db.Get(provider, "SELECT * FROM providers WHERE provider=$1 LIMIT 1", name); err != nil {
		return err
	}

	return nil
}
