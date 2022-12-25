package handlers

import (
	"context"
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
	ap := new(pgmodels.AuthProvider)
	pd := new(pgmodels.ProvidersData)
	token := new(pgmodels.AccessToken)

	user.CheckOnExistsWithoutPassword(a.DB, req.Email, provider)
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
