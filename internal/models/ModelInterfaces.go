package models

import "github.com/jmoiron/sqlx"

type RootModel interface {
	Create(unique string) (*sqlx.DB, error)
	CheckAndUpdateDb(unique string) (*sqlx.DB, bool)
	createSubDir(service, unique string) (*sqlx.DB, error)
}

type OAuthModel interface {
	RootModel
	GetTokenByStr(db *sqlx.DB, token string) (AccessToken, error)
}
