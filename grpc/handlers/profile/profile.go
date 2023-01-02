package profile

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/pavel-one/GoStarter/api"
	"github.com/pavel-one/GoStarter/grpc/internal/models"
)

type UserProfileService struct {
	api.UnimplementedProfileServiceServer
	DB *sqlx.DB
}

func NewUserProfileService(db *sqlx.DB) *UserProfileService {
	ps := new(UserProfileService)
	ps.DB = db
	return ps
}

func (u *UserProfileService) Create(ctx context.Context, req *api.ProfileRequest) (*api.ProfileResponse, error) {
	user := new(models.User)
	p := new(models.Profile)

	user.Find(u.DB, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	p.FindByUserUUID(u.DB, user.ID)
	if p.ID != 0 {
		return nil, errors.New("profile already exists")
	}

	setRaws(p, req)
	p.UserID = req.UserUUID
	if err := p.Create(u.DB); err != nil {
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
	user := new(models.User)
	p := new(models.Profile)
	user.Find(u.DB, req.UserUUID)

	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	p.FindByUserUUID(u.DB, user.ID)
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
	user := new(models.User)
	p := new(models.Profile)

	user.Find(u.DB, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	p.FindByUserUUID(u.DB, user.ID)
	if p.ID == 0 {
		return nil, errors.New("profile not found")
	}

	setRaws(p, req)

	if err := p.Update(u.DB); err != nil {
		return nil, err
	}

	p.FindByUserUUID(u.DB, user.ID)
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
	user := new(models.User)
	p := new(models.Profile)

	user.Find(u.DB, req.UserUUID)
	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	p.FindByUserUUID(u.DB, user.ID)
	if p.ID == 0 {
		return nil, errors.New("profile not found")
	}

	if err := p.Delete(u.DB); err != nil {
		return nil, err
	}

	return &api.ProfileDeleted{Message: "Profile deleted successfully"}, nil
}

func setRaws(p *models.Profile, req *api.ProfileRequest) {
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

func toResponse(p *models.Profile, data map[string]string) *api.ProfileResponse {
	return &api.ProfileResponse{
		FirstName:  p.FirstName.String,
		LastName:   p.LastName.String,
		Other_Data: data,
		Address:    p.Address.String,
	}
}
