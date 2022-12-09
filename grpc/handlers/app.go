package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/internal/appauth"
	"github.com/pavel-one/GoStarter/internal/models"
	"golang.org/x/crypto/bcrypt"
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
	user := new(models.User)
	token := new(models.AccessToken)

	user.FindByUniqueAndService(a.DB, req.Email, "app")
	if user.ID != 0 {
		fmt.Println("user already exists")
		return nil, errors.New("this user already exists")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashP)
	user.UniqueRaw = req.Email
	user.AuthorizedBy = "app"

	err = user.Create(a.DB)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, err
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

func (a *AppAuthService) Login(ctx context.Context, req *api.AppRequest) (*api.AppResponse, error) {
	user := new(models.User)
	token := new(models.AccessToken)

	if err := user.FindByUniqueAndService(a.DB, req.Email, "app"); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	tokenStr, err := appauth.GenerateToken(user.ID, user.UniqueRaw, "app")
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
