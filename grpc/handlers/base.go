package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/internal/pgmodels"
)

type BaseDB struct {
	DB *sqlx.DB
}

func ToAppResponse(user *pgmodels.User, token *pgmodels.AccessToken) *api.AppResponse {
	res := api.AppResponse{Struct: &api.User{}, TokenStr: ""}
	res.Struct.ID = uint64(user.ID)
	res.Struct.UniqueRaw = user.UniqueRaw
	res.Struct.AuthorizedBy = user.AuthorizedBy
	res.Struct.CreatedAt = user.CreatedAt.String()
	res.TokenStr = token.Token

	return &res
}
