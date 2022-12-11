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
	res := api.AppResponse{Struct: &api.User{}, TokenStr: ""}
	res.Struct.ID = uint64(user.ID)
	res.Struct.UniqueRaw = user.UniqueRaw
	res.Struct.AuthorizedBy = user.AuthorizedBy
	res.Struct.CreatedAt = user.CreatedAt.String()
	res.TokenStr = token.Token

	return &res
}
