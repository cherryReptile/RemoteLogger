package handlers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/internal/pgmodels"
	"github.com/pavel-one/GoStarter/internal/appauth"
)

type TelegramAuthService struct {
	api.UnimplementedAuthTelegramServiceServer
	BaseDB
}

func NewTelegramAuthService(db *sqlx.DB) *TelegramAuthService {
	gs := new(TelegramAuthService)
	gs.DB = db
	return gs
}

func (a *TelegramAuthService) Login(ctx context.Context, req *api.TelegramRequest) (*api.AppResponse, error) {
	user := new(pgmodels.User)
	token := new(pgmodels.AccessToken)

	user.FindByUniqueAndService(a.DB, req.Username, "telegram")
	if user.ID == 0 {
		user.UniqueRaw = req.Username
		user.AuthorizedBy = "telegram"
		if err := user.Create(a.DB); err != nil {
			return nil, err
		}
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
