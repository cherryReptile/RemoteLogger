package auth

import (
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/domain"
)

func ToAppResponse(user *domain.User, token *domain.AuthToken) *api.AppResponse {
	res := api.AppResponse{User: &api.User{}, JWTToken: ""}
	res.User.ID = user.ID
	res.User.Login = user.Login
	res.User.CreatedAt = user.CreatedAt.String()
	res.JWTToken = token.Token

	return &res
}

func ToAddedResponse(message string, user *domain.User) *api.AddedResponse {
	return &api.AddedResponse{
		Message: message,
		User: &api.User{
			ID:        user.ID,
			Login:     user.Login,
			CreatedAt: user.CreatedAt.String(),
		}}
}
