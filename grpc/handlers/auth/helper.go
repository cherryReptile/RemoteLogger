package auth

import (
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/domain"
)

func ToAppResponse(user *domain.User, token *domain.AuthToken) *api.AppResponse {
	res := api.AppResponse{Struct: &api.User{}, JWTToken: ""}
	res.Struct.ID = user.ID
	res.Struct.Login = user.Login
	res.Struct.CreatedAt = user.CreatedAt.String()
	res.JWTToken = token.Token

	return &res
}

func ToAddedResponse(message string, user *domain.User) *api.AddedResponse {
	return &api.AddedResponse{
		Message: message,
		Struct: &api.User{
			ID:        user.ID,
			Login:     user.Login,
			CreatedAt: user.CreatedAt.String(),
		}}
}
