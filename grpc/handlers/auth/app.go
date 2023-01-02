package auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/internal/authtoken"
	"github.com/pavel-one/GoStarter/grpc/internal/models"
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
	provider := "app"
	user := new(models.User)
	p := new(models.Provider)
	pd := new(models.ProvidersData)
	token := new(models.AccessToken)

	p.GetByProvider(a.DB, provider)
	if p.ID == 0 {
		return nil, errors.New("unknown auth provider")
	}

	pd.FindByUsernameAndProvider(a.DB, req.Email, p.ID)
	if pd.ID != 0 {
		return nil, errors.New("this user already exists")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Login = req.Email

	err = user.Create(a.DB, p.ID)
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, err
	}

	json, err := json.Marshal(map[string]string{"email": req.Email, "password": string(hashP)})
	if err != nil {
		return nil, err
	}

	pd.UserData = json
	pd.UserID = user.ID
	pd.ProviderID = p.ID
	pd.Username = user.Login
	if err = pd.Create(a.DB); err != nil {
		return nil, err
	}

	tokenStr, err := authtoken.GenerateToken(user.ID, user.Login, "app")
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
	userData := new(AppUserData)
	user := new(models.User)
	p := new(models.Provider)
	pd := new(models.ProvidersData)
	token := new(models.AccessToken)

	p.GetByProvider(a.DB, "app")
	if p.ID == 0 {
		return nil, errors.New("unknown provider")
	}

	pd.FindByUsernameAndProvider(a.DB, req.Email, p.ID)
	if pd.ID == 0 {
		return nil, errors.New("user not found")
	}

	user.Find(a.DB, pd.UserID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	if err := json.Unmarshal(pd.UserData, &userData); err != nil {
		return nil, err
	}

	if userData.Email == "" || userData.Password == "" {
		return nil, errors.New("email or password is required from db response")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	tokenStr, err := authtoken.GenerateToken(user.ID, user.Login, "app")
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
	user := new(models.User)
	up := new(models.UsersProviders)
	pd := new(models.ProvidersData)
	p := new(models.Provider)

	user.Find(a.DB, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("invalid user's uuid")
	}

	p.GetByProvider(a.DB, provider)
	if p.ID == 0 {
		return nil, errors.New("unknown provider")
	}

	pd.FindByUsernameAndProvider(a.DB, req.Request.Email, p.ID)
	if pd.ID != 0 {
		return nil, errors.New("user already exists")
	}

	if err := up.Create(a.DB, req.UserUUID, p.ID); err != nil {
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
	pd.ProviderID = p.ID
	pd.Username = req.Request.Email
	if err = pd.Create(a.DB); err != nil {
		return nil, err
	}

	return ToAddedResponse("App account added successfully", user), nil
}
