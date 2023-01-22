package repository

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/jmoiron/sqlx"
)

type clientUserRepository struct {
	db *sqlx.DB
}

func NewUserAndProfileRepository(db *sqlx.DB) domain.ClientUserRepo {
	return &clientUserRepository{
		db: db,
	}
}

func (r *clientUserRepository) GetUserWithProfile(clientUser *domain.ClientUser, userID string) error {
	if err := r.db.Get(clientUser,
		`select users.id, users.login, users.created_at, up.first_name, up.last_name, up.address, up.other_data 
	from users
    	left join user_profiles up on up.user_id = users.id
	where users.id = $1 limit 1`, userID); err != nil {
		return err
	}
	return nil
}

func (r *clientUserRepository) GetAuthClientUser(clientUser *domain.ClientUser, userID, token string) error {
	if err := r.db.Get(clientUser,
		`select users.id, users.login, users.created_at, a.token, up.first_name, up.last_name, up.address, up.other_data
	from users
	    left join access_tokens a on a.user_id = users.id
    	left join user_profiles up on up.user_id = users.id
	where users.id = $1 and a.token = $2 limit 1`, userID, token); err != nil {
		return err
	}
	return nil
}
