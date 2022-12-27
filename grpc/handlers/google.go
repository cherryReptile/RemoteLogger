package handlers

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/internal/appauth"
	"github.com/pavel-one/GoStarter/grpc/internal/pgmodels"
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
	provider := "google"
	user := new(pgmodels.User)
	ap := new(pgmodels.Provider)
	pd := new(pgmodels.ProvidersData)
	token := new(pgmodels.AccessToken)

	user.FindByLoginAndProvider(a.DB, req.Email, provider)
	if user.ID == "" {
		user.Login = req.Email
		if err := user.Create(a.DB, provider); err != nil {
			return nil, err
		}
	}

	if err := ap.GetByProvider(a.DB, provider); err != nil {
		return nil, err
	}

	pd.UserData = req.Data
	pd.UserID = user.ID
	pd.ProviderID = ap.ID
	pd.FindByUserUUIDAndProviderID(a.DB, user.ID, ap.ID)
	if pd.ID == 0 {
		if err := pd.Create(a.DB); err != nil {
			return nil, err
		}
	}

	tokenStr, err := appauth.GenerateToken(user.ID, user.Login, provider)
	if err != nil {
		return nil, err
	}

	token.Token = tokenStr
	token.UserUUID = user.ID
	if err = token.Create(a.DB); err != nil {
		return nil, err
	}

	return ToAppResponse(user, token), nil
}

func (a *GoogleAuthService) AddAccount(ctx context.Context, req *api.AddGoogleRequest) (*api.AddedResponse, error) {
	provider := "google"
	user := new(pgmodels.User)
	inter := new(pgmodels.Intermediate)
	pd := new(pgmodels.ProvidersData)
	ap := new(pgmodels.Provider)

	user.FindByLoginAndProvider(a.DB, req.Request.Email, provider)
	if user.ID != "" {
		return nil, errors.New("sorry this user authorized regardless of this account")
	}
	user.Find(a.DB, req.UserUUID)

	if err := ap.GetByProvider(a.DB, provider); err != nil {
		return nil, err
	}

	inter.Find(a.DB, req.UserUUID, ap.ID)
	if inter.ID != 0 {
		return nil, errors.New("sorry this account already been added")
	}

	if err := inter.Create(a.DB, req.UserUUID, ap.ID); err != nil {
		return nil, err
	}

	pd.UserData = req.Request.Data
	pd.UserID = req.UserUUID
	pd.ProviderID = ap.ID
	if err := pd.Create(a.DB); err != nil {
		return nil, err
	}

	return ToAddedResponse("Google account added successfully", user), nil
}
