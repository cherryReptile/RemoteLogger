package handlers

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/internal/pgmodels"
	"github.com/pavel-one/GoStarter/internal/appauth"
)

type GoogleAuthService struct {
	api.UnimplementedAuthGoogleServiceServer
	BaseDB
}

func NewGoogleAuthService(db *sqlx.DB) *GoogleAuthService {
	gs := new(GoogleAuthService)
	gs.DB = db
	return gs
}

func (a *GoogleAuthService) Login(ctx context.Context, req *api.GoogleRequest) (*api.AppResponse, error) {
	user := new(pgmodels.User)
	token := new(pgmodels.AccessToken)

	user.FindByUniqueAndService(a.DB, req.Email, "google")
	if user.ID == 0 {
		user.UniqueRaw = req.Email
		user.AuthorizedBy = "google"
		if err := user.Create(a.DB); err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	tokenStr, err := appauth.GenerateToken(user.ID, user.UniqueRaw, user.AuthorizedBy)
	if err != nil {
		return nil, err
	}

	token.Token = tokenStr
	token.UserID = user.ID
	if err = token.Create(a.DB); err != nil {
		return nil, err
	}

	return ToAppResponse(user, token), nil
}
