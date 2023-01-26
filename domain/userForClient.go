package domain

import "github.com/jmoiron/sqlx"

type ClientUser struct {
	User
	Profile
	AuthToken
}

type ClientUserRepo interface {
	GetUserWithProfile(clientUser *ClientUser, userID string) error
	GetAuthClientUser(clientUser *ClientUser, userID, token string) error
	GetAllWithOrderBy(field, orderBy string) (*sqlx.Rows, error)
	GetAllWithOrderByAndFilter(filter map[string]string, field, orderBy string) (*sqlx.Rows, error)
}

type ClientUserUsecase interface {
	GetUserWithProfile(userAndProfile *ClientUser, userID string) error
	GetAuthClientUser(clientUser *ClientUser, userID, token string) error
	GetAllWithOrderBy(field, orderBy string) (*sqlx.Rows, error)
	GetAllWithOrderByAndFilter(filter map[string]string, field, orderBy string) (*sqlx.Rows, error)
}
