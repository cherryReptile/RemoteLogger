package auth

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/repository"
	"github.com/cherryReptile/WS-AUTH/usecase"
	"github.com/jmoiron/sqlx"
)

type GitHubAuthService struct {
	api.UnimplementedAuthGithubServiceServer
	BaseOAuthHandler
}

func NewGitHubAuthService(db *sqlx.DB) *GitHubAuthService {
	gs := new(GitHubAuthService)
	gs.userUsecase = usecase.NewUserUsecase(repository.NewUserRepository(db))
	gs.providerUsecase = usecase.NewProviderUsecase(repository.NewProviderRepository(db))
	gs.tokenUsecase = usecase.NewTokenUsecase(repository.NewTokenRepository(db))
	gs.providersDataUsecase = usecase.NewProvidersDataUsecase(repository.NewProvidersDataRepo(db))
	gs.usersProvidersUsecase = usecase.NewUsersProvidersUsecase(repository.NewUsersProvidersRepository(db))
	gs.DB = db
	gs.Provider = "github"
	return gs
}

func (a *GitHubAuthService) Login(ctx context.Context, req *api.OAuthRequest) (*api.AppResponse, error) {
	user, token, err := a.LoginDefault(req)
	if err != nil {
		return nil, err
	}
	return ToAppResponse(user, token), nil
	return nil, nil
}

func (a *GitHubAuthService) AddAccount(ctx context.Context, req *api.AddOauthRequest) (*api.AddedResponse, error) {
	user, err := a.AddAccountDefault(req)
	if err != nil {
		return nil, err
	}
	return ToAddedResponse("GitHub account added successfully", user), nil
}
