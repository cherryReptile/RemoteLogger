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
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type appAuthService struct {
	api.UnimplementedAuthAppServiceServer
	BaseHandler
	DB *sqlx.DB
}

func NewAppAuthService(db *sqlx.DB) api.AuthAppServiceServer {
	as := new(appAuthService)
	as.userUsecase = usecase.NewUserUsecase(repository.NewUserRepository(db))
	as.providerUsecase = usecase.NewProviderUsecase(repository.NewProviderRepository(db))
	as.tokenUsecase = usecase.NewTokenUsecase(repository.NewTokenRepository(db))
	as.providersDataUsecase = usecase.NewProvidersDataUsecase(repository.NewProvidersDataRepo(db))
	as.usersProvidersUsecase = usecase.NewUsersProvidersUsecase(repository.NewUsersProvidersRepository(db))
	as.DB = db
	return as
}

func (s *appAuthService) Register(ctx context.Context, req *api.AppRequest) (*api.AppResponse, error) {
	provider := "app"
	var e struct {
		Email string `validate:"required,email"`
	}
	validate := validator.New()
	user := new(domain.User)
	p := new(domain.Provider)
	pd := new(domain.ProvidersData)
	up := new(domain.UsersProviders)
	token := new(domain.AuthToken)

	e.Email = req.Email
	if err := validate.Struct(e); err != nil {
		return nil, err
	}

	s.providerUsecase.GetByProvider(p, provider)
	if p.ID == 0 {
		return nil, errors.New("unknown auth provider")
	}

	s.providersDataUsecase.FindByUsernameAndProvider(pd, req.Email, p.ID)
	if pd.ID != 0 {
		return nil, errors.New("this user already exists")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Login = req.Email

	err = s.userUsecase.Create(user)
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, err
	}

	if err = s.usersProvidersUsecase.Create(up, user.ID, p.ID); err != nil {
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
	if err = s.providersDataUsecase.Create(pd); err != nil {
		return nil, err
	}

	tokenStr, err := authtoken.GenerateToken(user.ID, user.Login, "app")
	if err != nil {
		return nil, err
	}

	token.Token = tokenStr
	token.UserUUID = user.ID

	if err = s.tokenUsecase.Create(token); err != nil {
		return nil, err
	}

	return ToAppResponse(user, token), nil
}

func (s *appAuthService) Login(ctx context.Context, req *api.AppRequest) (*api.AppResponse, error) {
	userData := new(AppUserData)
	user := new(domain.User)
	p := new(domain.Provider)
	pd := new(domain.ProvidersData)
	token := new(domain.AuthToken)

	s.providerUsecase.GetByProvider(p, "app")
	if p.ID == 0 {
		return nil, errors.New("unknown provider")
	}

	s.providersDataUsecase.FindByUsernameAndProvider(pd, req.Email, p.ID)
	if pd.ID == 0 {
		return nil, errors.New("user not found")
	}

	s.userUsecase.Find(user, pd.UserID)
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
	if err = s.tokenUsecase.Create(token); err != nil {
		return nil, err
	}

	return ToAppResponse(user, token), nil
}

func (s *appAuthService) AddAccount(ctx context.Context, req *api.AddAppRequest) (*api.AddedResponse, error) {
	provider := "app"
	user := new(domain.User)
	up := new(domain.UsersProviders)
	pd := new(domain.ProvidersData)
	p := new(domain.Provider)

	s.userUsecase.Find(user, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("invalid user's uuid")
	}

	s.providerUsecase.GetByProvider(p, provider)
	if p.ID == 0 {
		return nil, errors.New("unknown provider")
	}

	s.providersDataUsecase.FindByUsernameAndProvider(pd, req.Request.Email, p.ID)
	if pd.ID != 0 {
		return nil, errors.New("user already exists")
	}

	pds, err := s.providersDataUsecase.GetAllByProvider(user.ID, p.ID)
	if err != nil {
		return nil, err
	}
	if len(pds) >= 1 {
		return nil, errors.New("you already have account in app")
	}

	if err = s.usersProvidersUsecase.Create(up, req.UserUUID, p.ID); err != nil {
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
	if err = s.providersDataUsecase.Create(pd); err != nil {
		return nil, err
	}

	return ToAddedResponse("Server account added successfully", user), nil
}
