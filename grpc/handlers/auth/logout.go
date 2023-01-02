package auth

import (
	"context"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/grpc/internal/models"
	"github.com/jmoiron/sqlx"
)

type LogoutService struct {
	api.UnimplementedLogoutServiceServer
	BaseDB
}

func NewLogoutAuthService(db *sqlx.DB) *LogoutService {
	ls := new(LogoutService)
	ls.DB = db
	return ls
}

func (l *LogoutService) Logout(ctx context.Context, req *api.TokenRequest) (*api.LogoutResponse, error) {
	token := new(models.AccessToken)

	if err := token.GetByToken(l.DB, req.Token); err != nil {
		return nil, err
	}

	if token.ID == 0 {
		return nil, errors.New("token not found")
	}

	if err := token.Delete(l.DB); err != nil {
		return nil, err
	}

	return &api.LogoutResponse{Message: "logout successfully"}, nil
}
