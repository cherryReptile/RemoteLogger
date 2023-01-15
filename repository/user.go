package repository

import (
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepo {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(user *domain.User) error {
	user.ID = uuid.NewString()
	user.CreatedAt = time.Now()

	_, err := u.db.NamedExec(`INSERT INTO users (id, login, created_at) 
								VALUES (:id, :login, :created_at)`, user)

	if err != nil {
		//return errors.New("failed to create user " + err.Error())
		return err
	}

	// update model
	if err = u.db.Get(user, "SELECT * FROM users WHERE id=$1 LIMIT 1", user.ID); err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Find(user *domain.User, uuid string) error {
	if err := u.db.Get(user, "SELECT * FROM users WHERE id=$1", uuid); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) FindByLoginAndProvider(user *domain.User, login, provider string) error {
	if err := u.db.Get(user,
		`select users.id, login, users.created_at 
	from users
    	left join users_providers up on up.user_id = users.id
    	left join providers p on up.provider_id = p.id
	where users.login = $1 and p.provider = $2 limit 1`, login, provider); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetTokenByStr(user *domain.User, tokenStr string) (*domain.AuthToken, error) {
	t := domain.AuthToken{}
	err := u.db.Get(&t, "SELECT * FROM access_tokens WHERE user_id=$1 AND token=$2 ORDER BY id DESC", user.ID, tokenStr)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
