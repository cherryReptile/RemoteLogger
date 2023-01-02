package auth

import (
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/grpc/internal/authtoken"
	"github.com/cherryReptile/WS-AUTH/grpc/internal/models"
	"github.com/jmoiron/sqlx"
)

type BaseDB struct {
	DB *sqlx.DB
}

type BaseOAuthHandler struct {
	BaseDB
	Provider string
}

func (h *BaseOAuthHandler) LoginDefault(req *api.OAuthRequest) (*models.User, *models.AccessToken, error) {
	user := new(models.User)
	token := new(models.AccessToken)
	p := new(models.Provider)
	pd := new(models.ProvidersData)

	if err := p.GetByProvider(h.DB, h.Provider); err != nil {
		return nil, nil, err
	}

	pd.FindByUsernameAndProvider(h.DB, req.Username, p.ID)
	if pd.ID == 0 {
		user.Login = req.Username
		if err := user.Create(h.DB, p.ID); err != nil {
			return nil, nil, err
		}
	}

	if user.ID == "" {
		user.Find(h.DB, pd.UserID)
		if user.ID == "" {
			return nil, nil, errors.New("user not found")
		}
	}

	if pd.ID == 0 {
		pd.UserData = req.Data
		pd.UserID = user.ID
		pd.ProviderID = p.ID
		pd.Username = user.Login
		if err := pd.Create(h.DB); err != nil {
			return nil, nil, err
		}
	}

	tokenStr, err := authtoken.GenerateToken(user.ID, user.Login, h.Provider)
	if err != nil {
		return nil, nil, err
	}

	token.Token = tokenStr
	token.UserUUID = user.ID
	if err = token.Create(h.DB); err != nil {
		return nil, nil, err
	}

	return user, token, nil
}

func (h *BaseOAuthHandler) AddAccountDefault(req *api.AddOauthRequest) (*models.User, error) {
	user := new(models.User)
	up := new(models.UsersProviders)
	pd := new(models.ProvidersData)
	p := new(models.Provider)

	user.Find(h.DB, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	p.GetByProvider(h.DB, h.Provider)
	if p.ID == 0 {
		return nil, errors.New("unknown provider")
	}

	pd.FindByUsernameAndProvider(h.DB, req.Request.Username, p.ID)
	if pd.ID != 0 {
		return nil, errors.New("user already exists")
	}

	if err := up.Create(h.DB, req.UserUUID, p.ID); err != nil {
		return nil, err
	}

	pd.UserData = req.Request.Data
	pd.UserID = req.UserUUID
	pd.ProviderID = p.ID
	pd.Username = req.Request.Username
	if err := pd.Create(h.DB); err != nil {
		return nil, err
	}

	return user, nil
}

func ToAppResponse(user *models.User, token *models.AccessToken) *api.AppResponse {
	res := api.AppResponse{Struct: &api.User{}, TokenStr: ""}
	res.Struct.UUID = user.ID
	res.Struct.Login = user.Login
	res.Struct.CreatedAt = user.CreatedAt.String()
	res.TokenStr = token.Token

	return &res
}

func ToAddedResponse(message string, user *models.User) *api.AddedResponse {
	return &api.AddedResponse{
		Message: message,
		Struct: &api.User{
			UUID:      user.ID,
			Login:     user.Login,
			CreatedAt: user.CreatedAt.String(),
		}}
}
