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
	ap := new(pgmodels.AuthProvider)
	pd := new(pgmodels.ProvidersData)

	user.CheckOnExistsWithoutPassword(a.DB, req.Login, provider)
	if user.ID == "" {
		user.Login = req.Login
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

func (a *GitHubAuthService) AddAccount(ctx context.Context, req *api.AddGitHubRequest) (*api.AddedResponse, error) {
	provider := "github"
	user := new(pgmodels.User)
	inter := new(pgmodels.Intermediate)
	pd := new(pgmodels.ProvidersData)
	ap := new(pgmodels.AuthProvider)

	user.CheckOnExistsWithoutPassword(a.DB, req.Request.Login, provider)
	if user.ID != "" {
		return nil, errors.New("sorry this user authorized regardless of this account")
	}

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

	return &api.AddedResponse{
		Message: "GitHub account added successfully",
		Struct: &api.User{
			UUID:      user.ID,
			Login:     user.Login,
			CreatedAt: user.CreatedAt.String(),
		}}, nil
}
