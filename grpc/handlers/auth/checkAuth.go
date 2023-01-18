package auth

import (
	"context"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/cherryReptile/WS-AUTH/internal/authtoken"
	"github.com/cherryReptile/WS-AUTH/repository"
	"github.com/cherryReptile/WS-AUTH/usecase"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

type checkAuthService struct {
	api.UnimplementedCheckAuthServiceServer
	userUsecase  domain.UserUsecase
	tokenUsecase domain.AuthTokenUsecase
	DB           *sqlx.DB
}

func NewCheckAuthService(db *sqlx.DB) api.CheckAuthServiceServer {
	cas := new(checkAuthService)
	cas.userUsecase = usecase.NewUserUsecase(repository.NewUserRepository(db))
	cas.tokenUsecase = usecase.NewTokenUsecase(repository.NewTokenRepository(db))
	cas.DB = db
	return cas
}

func (s *checkAuthService) CheckAuth(ctx context.Context, req *api.TokenRequest) (*api.CheckAuthResponse, error) {
	user := new(domain.User)
	token := new(domain.AuthToken)
	claims, err := authtoken.GetClaims(req.Token)
	if err != nil {
		err, ok := err.(*jwt.ValidationError)
		if !ok {
			return nil, err
		}

		if err.Errors == 16 {
			s.tokenUsecase.GetByToken(token, req.Token)
			if token.ID == 0 {
				return nil, err
			}

			s.tokenUsecase.Delete(token)
		}
		return nil, err
	}

	if err = s.userUsecase.FindByLoginAndProvider(user, claims.Unique, claims.Service); err != nil {
		return nil, errors.New("user not found")
	}

	token, err = s.userUsecase.GetTokenByStr(user, req.Token)
	if err != nil {
		return nil, errors.New("token not found")
	}
	if token.ID == 0 {
		return nil, err
	}

	return &api.CheckAuthResponse{UserUUID: user.ID}, nil
}
