package handlers

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
)

type AppAuthService struct {
	api.UnimplementedAuthAppServiceServer
	BaseDB
}

func NewAppAuthService(db *sqlx.DB) *AppAuthService {
	as := new(AppAuthService)
	as.DB = db
	return as
}

func (a *AppAuthService) Register(ctx context.Context, req *api.AppRequest) (*api.AppResponse, error) {
	//user := new(pgmodels.User)
	//token := new(pgmodels.AccessToken)
	//
	//user.CheckOnExistsWithoutPassword(a.DB, req.Email, "app")
	//if user.ID != "" {
	//	return nil, errors.New("this user already exists")
	//}
	//
	//hashP, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	return nil, err
	//}
	//
	//user.Password = string(hashP)
	//user.Login = req.Email
	//
	//err = user.Create(a.DB, "app")
	//if err != nil {
	//	return nil, err
	//}
	//
	//if user.ID == "" {
	//	return nil, err
	//}
	//
	//tokenStr, err := appauth.GenerateToken(user.ID, user.Login, "app")
	//if err != nil {
	//	return nil, err
	//}
	//
	//token.Token = tokenStr
	//token.UserUUID = user.ID
	//
	//if err = token.Create(a.DB); err != nil {
	//	return nil, err
	//}
	//
	//return ToAppResponse(user, token), nil
	return nil, nil
}

func (a *AppAuthService) Login(ctx context.Context, req *api.AppRequest) (*api.AppResponse, error) {
	//user := new(pgmodels.User)
	//token := new(pgmodels.AccessToken)
	//
	//if err := user.CheckOnExistsWithoutPassword(a.DB, req.Email, "app"); err != nil {
	//	return nil, err
	//}
	//
	//if user.ID == "" {
	//	return nil, errors.New("user not found")
	//}
	//
	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	//	return nil, err
	//}
	//
	//tokenStr, err := appauth.GenerateToken(user.ID, user.Login, "app")
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
