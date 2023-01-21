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
	//userUsecase           domain.UserUsecase
	clientUserUsecase domain.ClientUserUsecase
	tokenUsecase      domain.AuthTokenUsecase
	//profileUsecase        domain.ProfileUsecase
	DB *sqlx.DB
}

func NewGetUserService(db *sqlx.DB) api.GetUserServiceServer {
	cas := new(getUserService)
	cas.tokenUsecase = usecase.NewTokenUsecase(repository.NewTokenRepository(db))
	//cas.userUsecase = usecase.NewUserUsecase(repository.NewUserRepository(db))
	cas.clientUserUsecase = usecase.NewUserAndProfileUsecase(repository.NewUserAndProfileRepository(db))
	//cas.profileUsecase = usecase.NewProfileUsecase(repository.NewProfileRepository(db))
	cas.DB = db
	return cas
}

func (s *getUserService) GetUser(ctx context.Context, req *api.TokenRequest) (*api.UserClientResponse, error) {
	var od map[string]string
	clientUser := new(domain.ClientUser)
	clientUser.User = domain.User{}
	clientUser.Profile = domain.Profile{}
	clientUser.AuthToken = domain.AuthToken{}
	//profile := new(domain.Profile)
	claims, err := authtoken.GetClaims(req.Token)
	if err != nil {
		token := new(domain.AuthToken)
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

	//if err = s.userUsecase.FindByLoginAndProvider(user, claims.Unique, claims.Service); err != nil {
	//	return nil, errors.New("user not found")
	//}
	s.clientUserUsecase.GetAuthClientUser(clientUser, claims.UserID, req.Token)
	//s.userUsecase.GetUserWithProfile(clientUser, claims.UserID)
	if clientUser.User.ID == "" {
		return nil, errors.New("failed to get user")
	}

	//token, err = s.userUsecase.GetTokenByStr(&clientUser.User, req.Token)
	//if err != nil {
	//	return nil, errors.New("token not found")
	//}
	if clientUser.AuthToken.Token == "" {
		return nil, errors.New("failed to get token")
	}

	//if err = s.profileUsecase.FindByUserUUID(profile, user.ID); err != nil {
	//	return nil, err
	//}

	if clientUser.Profile.OtherData != nil {
		if err = json.Unmarshal(clientUser.Profile.OtherData, &od); err != nil {
			return nil, err
		}
	}

	return &api.UserClientResponse{
		User: &api.User{
			ID:        clientUser.User.ID,
			Login:     clientUser.User.Login,
			CreatedAt: clientUser.User.CreatedAt.String(),
		},
		Profile: &api.ProfileResponse{
			FirstName:  clientUser.Profile.FirstName.String,
			LastName:   clientUser.Profile.LastName.String,
			Address:    clientUser.Profile.Address.String,
			Other_Data: od,
		},
	}, nil
}
