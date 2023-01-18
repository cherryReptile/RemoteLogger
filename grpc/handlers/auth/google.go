package auth

import (
	"context"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/repository"
	"github.com/cherryReptile/WS-AUTH/usecase"
	"github.com/jmoiron/sqlx"
)

type googleAuthService struct {
	api.UnimplementedAuthGoogleServiceServer
	BaseOAuthHandler
}

func NewGoogleAuthService(db *sqlx.DB) api.AuthGoogleServiceServer {
	gs := new(googleAuthService)
	gs.userUsecase = usecase.NewUserUsecase(repository.NewUserRepository(db))
	gs.providerUsecase = usecase.NewProviderUsecase(repository.NewProviderRepository(db))
	gs.tokenUsecase = usecase.NewTokenUsecase(repository.NewTokenRepository(db))
	gs.providersDataUsecase = usecase.NewProvidersDataUsecase(repository.NewProvidersDataRepo(db))
	gs.usersProvidersUsecase = usecase.NewUsersProvidersUsecase(repository.NewUsersProvidersRepository(db))
	gs.DB = db
	gs.Provider = "google"
	return gs
}

func (s *googleAuthService) Login(ctx context.Context, req *api.OAuthRequest) (*api.AppResponse, error) {
	user, token, err := s.LoginDefault(req)
	if err != nil {
		return nil, err
	}
	return ToAppResponse(user, token), nil
}

func (s *googleAuthService) AddAccount(ctx context.Context, req *api.AddOauthRequest) (*api.AddedResponse, error) {
	user, err := s.AddAccountDefault(req)
	if err != nil {
		return nil, err
	}
	return ToAddedResponse("Google account added successfully", user), nil
}
