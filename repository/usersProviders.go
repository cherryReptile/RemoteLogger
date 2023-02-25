package repository

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/jmoiron/sqlx"
	"time"
)

type usersProvidersRepository struct {
	db *sqlx.DB
}

func NewUsersProvidersRepository(db *sqlx.DB) domain.UsersProvidersRepo {
	return &usersProvidersRepository{
		db: db,
	}
}

func (r *usersProvidersRepository) Create(up *domain.UsersProviders, userUUID string, providerID uint, tx *sqlx.Tx) error {
	create := `INSERT INTO users_providers (user_id, provider_id, created_at) 
								VALUES (:user_id, :provider_id, :created_at)`
	get := "SELECT * FROM users_providers ORDER BY id DESC LIMIT 1"
	up.UserID = userUUID
	up.ProviderID = providerID
	up.CreatedAt = time.Now()
	var err error

	if tx != nil {
		_, err = tx.NamedExec(create, up)

		if err != nil {
			return err
		}

		return tx.Get(up, get)
	}

	_, err = r.db.NamedExec(create, up)

	if err != nil {
		return err
	}

	return r.db.Get(up, get)
}
