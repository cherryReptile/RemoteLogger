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

type AuthService struct {
	api.UnimplementedAuthServiceServer
	DB *sqlx.DB
}

func NewAuthService(db *sqlx.DB) *AuthService {
	return &AuthService{DB: db}
}

func (a *AuthService) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisteredResponse, error) {
	user := new(models.User)
	tokenModel := new(models.AccessToken)

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

	tokenModel.Token = tokenStr
	tokenModel.UserID = user.ID

	if err = tokenModel.Create(a.DB); err != nil {
		return nil, err
	}

	return &api.RegisteredResponse{
		Usr: &api.User{
			ID:           uint64(user.ID),
			UniqueRaw:    user.UniqueRaw,
			Password:     user.Password,
			AuthorizedBy: user.AuthorizedBy,
		},
		TokenStr: tokenModel.Token,
	}, nil
}
