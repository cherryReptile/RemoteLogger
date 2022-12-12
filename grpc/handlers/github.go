package handlers

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/internal/appauth"
	"github.com/pavel-one/GoStarter/grpc/internal/pgmodels"
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
	user := new(pgmodels.User)
	token := new(pgmodels.AccessToken)

	user.FindByUniqueAndService(a.DB, req.Login, "github")
	if user.ID == 0 {
		user.UniqueRaw = req.Login
		user.AuthorizedBy = "github"
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
