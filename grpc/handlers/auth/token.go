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

type JWTTokenService struct {
	api.UnimplementedJWTTokenServiceServer
	tokenUsecase domain.AuthTokenUsecase
	DB           *sqlx.DB
}

func NewJWTTokenService(db *sqlx.DB) api.JWTTokenServiceServer {
	ls := new(JWTTokenService)
	ls.tokenUsecase = usecase.NewTokenUsecase(repository.NewTokenRepository(db))
	ls.DB = db
	return ls
}

func (s *JWTTokenService) Drop(ctx context.Context, req *api.TokenRequest) (*api.DroppedTokenResponse, error) {
	token := new(domain.AuthToken)

	if err := s.tokenUsecase.GetByToken(token, req.Token); err != nil {
		return nil, err
	}

	if token.ID == 0 {
		return nil, errors.New("token not found")
	}

	if err := s.tokenUsecase.Delete(token); err != nil {
		return nil, err
	}

	return &api.DroppedTokenResponse{Message: "token has been dropped successfully"}, nil
}
