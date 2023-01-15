package auth

import (
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/cherryReptile/WS-AUTH/grpc/internal/authtoken"
	"github.com/jmoiron/sqlx"
)

type BaseDB struct {
	DB *sqlx.DB
}

type BaseHandler struct {
	userUsecase           domain.UserUsecase
	providerUsecase       domain.ProviderUsecase
	tokenUsecase          domain.AuthTokenUsecase
	providersDataUsecase  domain.ProvidersDataUsecase
	usersProvidersUsecase domain.UsersProvidersUsecase
}

type BaseOAuthHandler struct {
	BaseDB
	BaseHandler
	Provider string
}

func (h *BaseOAuthHandler) LoginDefault(req *api.OAuthRequest) (*domain.User, *domain.AuthToken, error) {
	user := new(domain.User)
	token := new(domain.AuthToken)
	p := new(domain.Provider)
	up := new(domain.UsersProviders)
	pd := new(domain.ProvidersData)

	if err := h.providerUsecase.GetByProvider(p, h.Provider); err != nil {
		return nil, nil, err
	}

	h.providersDataUsecase.FindByUsernameAndProvider(pd, req.Username, p.ID)
	if pd.ID == 0 {
		user.Login = req.Username
		if err := h.userUsecase.Create(user); err != nil {
			return nil, nil, err
		}
		if err := h.usersProvidersUsecase.Create(up, user.ID, p.ID); err != nil {
			return nil, nil, err
		}
	}

	if user.ID == "" {
		h.userUsecase.Find(user, pd.UserID)
		if user.ID == "" {
			return nil, nil, errors.New("user not found")
		}
	}

	if pd.ID == 0 {
		pd.UserData = req.Data
		pd.UserID = user.ID
		pd.ProviderID = p.ID
		pd.Username = user.Login
		if err := h.providersDataUsecase.Create(pd); err != nil {
			return nil, nil, err
		}
	}

	tokenStr, err := authtoken.GenerateToken(user.ID, user.Login, h.Provider)
	if err != nil {
		return nil, nil, err
	}

	token.Token = tokenStr
	token.UserUUID = user.ID
	if err = h.tokenUsecase.Create(token); err != nil {
		return nil, nil, err
	}

	return user, token, nil
}

func (h *BaseOAuthHandler) AddAccountDefault(req *api.AddOauthRequest) (*domain.User, error) {
	user := new(domain.User)
	up := new(domain.UsersProviders)
	pd := new(domain.ProvidersData)
	p := new(domain.Provider)

	h.userUsecase.Find(user, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	h.providerUsecase.GetByProvider(p, h.Provider)
	if p.ID == 0 {
		return nil, errors.New("unknown provider")
	}

	h.providersDataUsecase.FindByUsernameAndProvider(pd, req.Request.Username, p.ID)
	if pd.ID != 0 {
		return nil, errors.New("user already exists")
	}

	if err := h.usersProvidersUsecase.Create(up, req.UserUUID, p.ID); err != nil {
		return nil, err
	}

	pd.UserData = req.Request.Data
	pd.UserID = req.UserUUID
	pd.ProviderID = p.ID
	pd.Username = req.Request.Username
	if err := h.providersDataUsecase.Create(pd); err != nil {
		return nil, err
	}

	return user, nil
}

func ToAppResponse(user *domain.User, token *domain.AuthToken) *api.AppResponse {
	res := api.AppResponse{Struct: &api.User{}, TokenStr: ""}
	res.Struct.UUID = user.ID
	res.Struct.Login = user.Login
	res.Struct.CreatedAt = user.CreatedAt.String()
	res.TokenStr = token.Token

	return &res
}

func ToAddedResponse(message string, user *domain.User) *api.AddedResponse {
	return &api.AddedResponse{
		Message: message,
		Struct: &api.User{
			UUID:      user.ID,
			Login:     user.Login,
			CreatedAt: user.CreatedAt.String(),
		}}
}
