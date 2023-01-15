package profile

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/cherryReptile/WS-AUTH/repository"
	"github.com/cherryReptile/WS-AUTH/usecase"
	"github.com/jmoiron/sqlx"
)

type UserProfileService struct {
	api.UnimplementedProfileServiceServer
	userUsecase    domain.UserUsecase
	profileUsecase domain.ProfileUsecase
	DB             *sqlx.DB
}

func NewUserProfileService(db *sqlx.DB) *UserProfileService {
	ps := new(UserProfileService)
	ps.userUsecase = usecase.NewUserUsecase(repository.NewUserRepository(db))
	ps.profileUsecase = usecase.NewProfileUsecase(repository.NewProfileRepository(db))
	ps.DB = db
	return ps
}

func (u *UserProfileService) Create(ctx context.Context, req *api.ProfileRequest) (*api.ProfileResponse, error) {
	user := new(domain.User)
	p := new(domain.Profile)

	u.userUsecase.Find(user, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	u.profileUsecase.FindByUserUUID(p, user.ID)
	if p.ID != 0 {
		return nil, errors.New("profile already exists")
	}

	setRaws(p, req)
	p.UserID = req.UserUUID
	if err := u.profileUsecase.Create(p); err != nil {
		return nil, err
	}

	var data map[string]string
	if err := json.Unmarshal(p.OtherData, &data); err != nil {
		return nil, err
	}

	return &api.ProfileResponse{
		FirstName:  p.FirstName.String,
		LastName:   p.LastName.String,
		Address:    p.Address.String,
		Other_Data: data,
	}, nil
}

func (u *UserProfileService) Get(ctx context.Context, req *api.ProfileUUID) (*api.ProfileResponse, error) {
	user := new(domain.User)
	p := new(domain.Profile)
	u.userUsecase.Find(user, req.UserUUID)

	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	u.profileUsecase.FindByUserUUID(p, user.ID)
	if p.ID == 0 {
		return nil, errors.New("profile not found")
	}

	var data map[string]string
	if err := json.Unmarshal(p.OtherData, &data); err != nil {
		return nil, err
	}

	return toResponse(p, data), nil
}

func (u *UserProfileService) Update(ctx context.Context, req *api.ProfileRequest) (*api.ProfileResponse, error) {
	user := new(domain.User)
	p := new(domain.Profile)

	u.userUsecase.Find(user, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	u.profileUsecase.FindByUserUUID(p, user.ID)
	if p.ID == 0 {
		return nil, errors.New("profile not found")
	}

	setRaws(p, req)

	if err := u.profileUsecase.Update(p); err != nil {
		return nil, err
	}

	u.profileUsecase.FindByUserUUID(p, user.ID)
	if p.ID == 0 {
		return nil, errors.New("failed to get user's profile after update")
	}

	var data map[string]string
	if err := json.Unmarshal(p.OtherData, &data); err != nil {
		return nil, err
	}

	return toResponse(p, data), nil
}

func (u *UserProfileService) Delete(ctx context.Context, req *api.ProfileUUID) (*api.ProfileDeleted, error) {
	user := new(domain.User)
	p := new(domain.Profile)

	u.userUsecase.Find(user, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	u.profileUsecase.FindByUserUUID(p, user.ID)
	if p.ID == 0 {
		return nil, errors.New("profile not found")
	}

	if err := u.profileUsecase.Delete(p); err != nil {
		return nil, err
	}

	return &api.ProfileDeleted{Message: "Profile deleted successfully"}, nil
}

func setRaws(p *domain.Profile, req *api.ProfileRequest) {
	switch {
	case req.FirstName != "":
		p.FirstName.String = req.FirstName
		p.FirstName.Valid = true
		fallthrough
	case req.LastName != "":
		p.LastName.String = req.LastName
		p.LastName.Valid = true
		fallthrough
	case req.Other_Data != nil:
		p.OtherData = req.Other_Data
		fallthrough
	case req.Address != "":
		p.Address.String = req.Address
		p.Address.Valid = true
	}
}

func toResponse(p *domain.Profile, data map[string]string) *api.ProfileResponse {
	return &api.ProfileResponse{
		FirstName:  p.FirstName.String,
		LastName:   p.LastName.String,
		Other_Data: data,
		Address:    p.Address.String,
	}
}
