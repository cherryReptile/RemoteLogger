package handlers

import (
	"context"
	"github.com/pavel-one/GoStarter/grpc/internal/appauth"
	"github.com/pavel-one/GoStarter/grpc/internal/pgmodels"

	//"errors"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	//"github.com/pavel-one/GoStarter/grpc/internal/appauth"
	//"github.com/pavel-one/GoStarter/grpc/internal/pgmodels"
)

type GitHubAuthService struct {
	api.UnimplementedAuthGithubServiceServer
	BaseDB
}

func NewGitHubAuthService(db *sqlx.DB) *GitHubAuthService {
	gs := new(GitHubAuthService)
	gs.DB = db
	return gs
}

func (a *GitHubAuthService) Login(ctx context.Context, req *api.GitHubRequest) (*api.AppResponse, error) {
	provider := "github"
	user := new(pgmodels.User)
	token := new(pgmodels.AccessToken)

	user.CheckOnExistsWithoutPassword(a.DB, req.Login, provider)
	if user.ID == "" {
		user.Login = req.Login
		if err := user.Create(a.DB, provider); err != nil {
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
