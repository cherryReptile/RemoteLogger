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
	res.Struct.UUID = user.ID
	res.Struct.Login = user.Login
	res.Struct.CreatedAt = user.CreatedAt.String()
	res.TokenStr = token.Token

	return &res
}

func ToAddedResponse(message string, user *pgmodels.User) *api.AddedResponse {
	return &api.AddedResponse{
		Message: message,
		Struct: &api.User{
			UUID:      user.ID,
			Login:     user.Login,
			CreatedAt: user.CreatedAt.String(),
		}}
}
