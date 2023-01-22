package auth

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/repository"
	"github.com/cherryReptile/WS-AUTH/usecase"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"os"
)

type gitHubAuthService struct {
	api.UnimplementedAuthGithubServiceServer
	BaseOAuthHandler
}

func NewGitHubAuthService(db *sqlx.DB) api.AuthGithubServiceServer {
	gs := new(gitHubAuthService)
	gs.userUsecase = usecase.NewUserUsecase(repository.NewUserRepository(db))
	gs.providerUsecase = usecase.NewProviderUsecase(repository.NewProviderRepository(db))
	gs.tokenUsecase = usecase.NewTokenUsecase(repository.NewTokenRepository(db))
	gs.providersDataUsecase = usecase.NewProvidersDataUsecase(repository.NewProvidersDataRepo(db))
	gs.usersProvidersUsecase = usecase.NewUsersProvidersUsecase(repository.NewUsersProvidersRepository(db))
	gs.profileUsecase = usecase.NewProfileUsecase(repository.NewProfileRepository(db))
	gs.DB = db
	gs.Config = &oauth2.Config{}
	gs.Config.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	gs.Config.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	gs.Config.Endpoint = github.Endpoint
	gs.Provider = "github"
	return gs
}

func (s *gitHubAuthService) GetToken(ctx context.Context, req *api.OAuthCodeRequest) (*api.OAuthTokenResponse, error) {
	return s.GetTokenDefault(req)
}

func (s *gitHubAuthService) Login(ctx context.Context, req *api.OAuthRequest) (*api.AppResponse, error) {
	user, token, err := s.LoginDefault(req)
	if err != nil {
		return nil, err
	}
	return ToAppResponse(user, token), nil
}

func (s *gitHubAuthService) AddAccount(ctx context.Context, req *api.AddOauthRequest) (*api.AddedResponse, error) {
	user, err := s.AddAccountDefault(req)
	if err != nil {
		return nil, err
	}
	return ToAddedResponse("GitHub account added successfully", user), nil
}
