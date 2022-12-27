package handlers

import (
	"context"
	"errors"
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
	p := new(pgmodels.Provider)
	pd := new(pgmodels.ProvidersData)

	if err := p.GetByProvider(a.DB, provider); err != nil {
		return nil, err
	}

	pd.FindByUsernameAndProvider(a.DB, req.Login, p.ID)
	if pd.ID == 0 {
		user.Login = req.Login
		if err := user.Create(a.DB, p.ID); err != nil {
			return nil, err
		}
	}

	if user.ID == "" {
		user.Find(a.DB, pd.UserID)
		if user.ID == "" {
			return nil, errors.New("user not found")
		}
	}

	if pd.ID == 0 {
		pd.UserData = req.Data
		pd.UserID = user.ID
		pd.ProviderID = p.ID
		pd.Username = user.Login
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

func (a *GitHubAuthService) AddAccount(ctx context.Context, req *api.AddGitHubRequest) (*api.AddedResponse, error) {
	provider := "github"
	user := new(pgmodels.User)
	up := new(pgmodels.UsersProviders)
	pd := new(pgmodels.ProvidersData)
	p := new(pgmodels.Provider)

	user.Find(a.DB, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	p.GetByProvider(a.DB, provider)
	if p.ID == 0 {
		return nil, errors.New("unknown provider")
	}

	pd.FindByUsernameAndProvider(a.DB, req.Request.Login, p.ID)
	if pd.ID != 0 {
		return nil, errors.New("user already exists")
	}

	if err := up.Create(a.DB, req.UserUUID, p.ID); err != nil {
		return nil, err
	}

	pd.UserData = req.Request.Data
	pd.UserID = req.UserUUID
	pd.ProviderID = p.ID
	pd.Username = req.Request.Login
	if err := pd.Create(a.DB); err != nil {
		return nil, err
	}

	return ToAddedResponse("GitHub account added successfully", user), nil
}
