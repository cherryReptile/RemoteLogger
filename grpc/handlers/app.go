package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/internal/appauth"
	"github.com/pavel-one/GoStarter/grpc/internal/pgmodels"
	"golang.org/x/crypto/bcrypt"
)

type AppAuthService struct {
	api.UnimplementedAuthAppServiceServer
	BaseDB
}

func NewAppAuthService(db *sqlx.DB) *AppAuthService {
	as := new(AppAuthService)
	as.DB = db
	return as
}

func (a *AppAuthService) Register(ctx context.Context, req *api.AppRequest) (*api.AppResponse, error) {
	user := new(pgmodels.User)
	ap := new(pgmodels.AuthProvider)
	pd := new(pgmodels.ProvidersData)
	token := new(pgmodels.AccessToken)

	user.FindByLoginAndProvider(a.DB, req.Email, "app")
	if user.ID != "" {
		return nil, errors.New("this user already exists")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Login = req.Email

	err = user.Create(a.DB, "app")
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, err
	}

	ap.GetByProvider(a.DB, "app")
	if ap.ID == 0 {
		return nil, errors.New("unknown auth provider")
	}

	json, err := json.Marshal(map[string]string{"password": string(hashP)})
	if err != nil {
		return nil, err
	}

	pd.UserData = json
	pd.UserID = user.ID
	pd.ProviderID = ap.ID
	if err = pd.Create(a.DB); err != nil {
		return nil, err
	}

	tokenStr, err := appauth.GenerateToken(user.ID, user.Login, "app")
	if err != nil {
		return nil, err
	}

	token.Token = tokenStr
	token.UserUUID = user.ID

	if err = token.Create(a.DB); err != nil {
		return nil, err
	}

	return ToAppResponse(user, token), nil
}

func (a *AppAuthService) Login(ctx context.Context, req *api.AppRequest) (*api.AppResponse, error) {
	user := new(pgmodels.User)
	ap := new(pgmodels.AuthProvider)
	pd := new(pgmodels.ProvidersData)
	token := new(pgmodels.AccessToken)

	if err := user.FindByLoginAndProvider(a.DB, req.Email, "app"); err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	ap.GetByProvider(a.DB, "app")
	if ap.ID == 0 {
		return nil, errors.New("unknown provider")
	}

	pd.FindByUserUUIDAndProviderID(a.DB, user.ID, ap.ID)
	if pd.ID == 0 {
		return nil, errors.New("user's provider data not found")
	}

	var data map[string]string
	err := json.Unmarshal(pd.UserData, &data)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data["password"]), []byte(req.Password)); err != nil {
		return nil, err
	}

	tokenStr, err := appauth.GenerateToken(user.ID, user.Login, "app")
	if err != nil {
		return nil, err
	}

	token.Token = tokenStr
	token.UserUUID = user.ID
	if err = token.Create(a.DB); err != nil {
		return nil, err
	}

	return ToAppResponse(user, token), nil
}

func (a *AppAuthService) AddAccount(ctx context.Context, req *api.AddAppRequest) (*api.AddedResponse, error) {
	provider := "app"
	user := new(pgmodels.User)
	inter := new(pgmodels.Intermediate)
	pd := new(pgmodels.ProvidersData)
	ap := new(pgmodels.AuthProvider)

	user.FindByLoginAndProvider(a.DB, req.Request.Email, provider)
	if user.ID != "" {
		return nil, errors.New("sorry this user authorized regardless of this account")
	}
	user.Find(a.DB, req.UserUUID)

	if err := ap.GetByProvider(a.DB, provider); err != nil {
		return nil, err
	}

	inter.Find(a.DB, req.UserUUID, ap.ID)
	if inter.ID != 0 {
		return nil, errors.New("sorry this account already been added")
	}

	if err := inter.Create(a.DB, req.UserUUID, ap.ID); err != nil {
		return nil, err
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(req.Request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	json, err := json.Marshal(map[string]string{"email": req.Request.Email, "password": string(hashP)})
	if err != nil {
		return nil, err
	}

	pd.UserData = json
	pd.UserID = req.UserUUID
	pd.ProviderID = ap.ID
	if err = pd.Create(a.DB); err != nil {
		return nil, err
	}

	return ToAddedResponse("App account added successfully", user), nil
}
