package handlers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
)

type GoogleAuthService struct {
	api.UnimplementedAuthGoogleServiceServer
	BaseOAuthHandler
}

func NewGoogleAuthService(db *sqlx.DB) *GoogleAuthService {
	gs := new(GoogleAuthService)
	gs.DB = db
	gs.Provider = "google"
	return gs
}

func (a *GoogleAuthService) Login(ctx context.Context, req *api.OAuthRequest) (*api.AppResponse, error) {
	user, token, err := a.LoginDefault(req)
	if err != nil {
		return nil, err
	}
	return ToAppResponse(user, token), nil
}

func (a *GoogleAuthService) AddAccount(ctx context.Context, req *api.AddOauthRequest) (*api.AddedResponse, error) {
	user, err := a.AddAccountDefault(req)
	if err != nil {
		return nil, err
	}
	return ToAddedResponse("Google account added successfully", user), nil
}
