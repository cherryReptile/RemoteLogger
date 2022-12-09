package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/internal/models"
)

type BaseDB struct {
	DB *sqlx.DB
}

func ToAppResponse(user *models.User, token *models.AccessToken) *api.AppResponse {
	res := new(api.AppResponse)
	res.Struct.ID = uint64(user.ID)
	res.Struct.UniqueRaw = user.UniqueRaw
	res.Struct.AuthorizedBy = user.AuthorizedBy
	res.Struct.CreatedAt = user.CreatedAt.String()
	res.TokenStr = token.Token
	if user.AuthorizedBy == "app" {
		res.Struct.Password = user.Password
	}

	return res
}
