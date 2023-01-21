package auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/cherryReptile/WS-AUTH/internal/authtoken"
	"github.com/cherryReptile/WS-AUTH/repository"
	"github.com/cherryReptile/WS-AUTH/usecase"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

type getUserService struct {
	api.UnimplementedGetUserServiceServer
	userUsecase    domain.UserUsecase
	tokenUsecase   domain.AuthTokenUsecase
	profileUsecase domain.ProfileUsecase
	DB             *sqlx.DB
}

func NewGetUserService(db *sqlx.DB) api.GetUserServiceServer {
	cas := new(getUserService)
	cas.userUsecase = usecase.NewUserUsecase(repository.NewUserRepository(db))
	cas.tokenUsecase = usecase.NewTokenUsecase(repository.NewTokenRepository(db))
	cas.profileUsecase = usecase.NewProfileUsecase(repository.NewProfileRepository(db))
	cas.DB = db
	return cas
}

func (s *getUserService) GetUser(ctx context.Context, req *api.TokenRequest) (*api.UserClientResponse, error) {
	var od map[string]string
	user := new(domain.User)
	token := new(domain.AuthToken)
	profile := new(domain.Profile)
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

	if err = s.profileUsecase.FindByUserUUID(profile, user.ID); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(profile.OtherData, &od); err != nil {
		return nil, err
	}

	return &api.UserClientResponse{
		User: &api.User{
			ID:        user.ID,
			Login:     user.Login,
			CreatedAt: user.CreatedAt.String(),
		},
		Profile: &api.ProfileResponse{
			FirstName:  profile.FirstName.String,
			LastName:   profile.LastName.String,
			Address:    profile.Address.String,
			Other_Data: od,
		},
	}, nil
}
