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

func (u *usersProvidersRepository) Create(up *domain.UsersProviders, userUUID string, providerID uint) error {
	up.UserID = userUUID
	up.ProviderID = providerID
	up.CreatedAt = time.Now()
	_, err := u.db.NamedExec(`INSERT INTO users_providers (user_id, provider_id, created_at) 
								VALUES (:user_id, :provider_id, :created_at)`, up)

	if err != nil {
		return err
	}

	// update model
	if err = u.db.Get(up, "SELECT * FROM users_providers ORDER BY id DESC LIMIT 1"); err != nil {
		return err
	}

	return nil
}
