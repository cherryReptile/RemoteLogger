package auth

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/jmoiron/sqlx"
)

type TelegramAuthService struct {
	api.UnimplementedAuthTelegramServiceServer
	DB *sqlx.DB
}

func NewTelegramAuthService(db *sqlx.DB) *TelegramAuthService {
	gs := new(TelegramAuthService)
	gs.DB = db
	return gs
}

func (a *TelegramAuthService) Login(ctx context.Context, req *api.TelegramRequest) (*api.AppResponse, error) {
	//user := new(pgmodels.User)
	//token := new(pgmodels.AccessToken)
	//
	//user.FindByUniqueAndService(a.DB, req.Username, "telegram")
	//if user.ID == 0 {
	//	user.UniqueRaw = req.Username
	//	user.AuthorizedBy = "telegram"
	//	if err := user.Create(a.DB); err != nil {
	//		return nil, err
	//	}
	//}
	//
	//tokenStr, err := appauth.GenerateToken(user.ID, user.UniqueRaw, user.AuthorizedBy)
	//if err != nil {
	//	return nil, err
	//}
	//
	//token.Token = tokenStr
	//token.UserUUID = user.ID
	//if err = token.Create(a.DB); err != nil {
	//	return nil, err
	//}
	//
	//return ToAppResponse(user, token), nil
	return nil, nil
}
