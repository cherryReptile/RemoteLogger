package models

import "github.com/jmoiron/sqlx"

type RootModel interface {
	Create(unique string) (*sqlx.DB, error)
	//CheckAndUpdateDb(unique string) (*sqlx.DB, bool)
	//createSubDir(service, unique string) (*sqlx.DB, error)
}

//type AuthModel interface {
//	RootModel
//	GetTokenByStr(db *sqlx.DB, token string) (AccessToken, error)
//	GetUniqueRaw() string
//}

//type JwtAuthModel interface {
//	OAuthModel
//	GetUniqueRaw() string
//}
