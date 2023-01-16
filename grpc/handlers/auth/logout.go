package auth

import (
	"context"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/cherryReptile/WS-AUTH/repository"
	"github.com/cherryReptile/WS-AUTH/usecase"
	"github.com/jmoiron/sqlx"
)

type logoutService struct {
	api.UnimplementedLogoutServiceServer
	tokenUsecase domain.AuthTokenUsecase
	DB           *sqlx.DB
}

func NewLogoutAuthService(db *sqlx.DB) api.LogoutServiceServer {
	ls := new(logoutService)
	ls.tokenUsecase = usecase.NewTokenUsecase(repository.NewTokenRepository(db))
	ls.DB = db
	return ls
}

func (l *logoutService) Logout(ctx context.Context, req *api.TokenRequest) (*api.LogoutResponse, error) {
	token := new(domain.AuthToken)

	if err := l.tokenUsecase.GetByToken(token, req.Token); err != nil {
		return nil, err
	}

	if token.ID == 0 {
		return nil, errors.New("token not found")
	}

	if err := l.tokenUsecase.Delete(token); err != nil {
		return nil, err
	}

	return &api.LogoutResponse{Message: "logout successfully"}, nil
}
