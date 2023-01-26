package userInfo

import (
	"encoding/json"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/cherryReptile/WS-AUTH/domain"
	"github.com/cherryReptile/WS-AUTH/repository"
	"github.com/cherryReptile/WS-AUTH/usecase"
	"github.com/jmoiron/sqlx"
)

type userInfoService struct {
	api.UnimplementedUserInfoServiceServer
	DB                *sqlx.DB
	clientUserUsecase domain.ClientUserUsecase
}

func NewUserInfoService(db *sqlx.DB) api.UserInfoServiceServer {
	us := new(userInfoService)
	us.DB = db
	us.clientUserUsecase = usecase.NewUserAndProfileUsecase(repository.NewUserAndProfileRepository(db))
	return us
}

func (s *userInfoService) GetAllUsersWithSortAndFilter(req *api.GetUsersRequest, stream api.UserInfoService_GetAllUsersWithSortAndFilterServer) error {
	var err error
	var rows *sqlx.Rows
	switch req.OrderBy {
	case "desc":
	case "asc":
	default:
		return errors.New("unknown param for sort")
	}
	switch req.Field {
	case "id":
	case "login":
	case "created_at":
	case "first_name":
	case "last_name":
	case "address":
	default:
		return errors.New("unsupportable field for sorting")
	}

	if req.FieldsAndValuesFilter == nil {
		rows, err = s.clientUserUsecase.GetAllWithOrderBy(req.Field, req.OrderBy)
	} else if len(req.FieldsAndValuesFilter) > 0 {
		rows, err = s.clientUserUsecase.GetAllWithOrderByAndFilter(req.FieldsAndValuesFilter, req.Field, req.OrderBy)
	}
	if err != nil {
		return err
	}

	i := 0
	for rows.Next() {
		var od map[string]string
		clientUser := new(domain.ClientUser)
		clientUser.User = domain.User{}
		clientUser.Profile = domain.Profile{}
		clientUser.AuthToken = domain.AuthToken{}
		if err = rows.StructScan(clientUser); err != nil {
			return err
		}

		if err = json.Unmarshal(clientUser.OtherData, &od); err != nil {
			return err
		}

		if err = stream.Send(&api.UserClientResponse{
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
			JWTToken: clientUser.AuthToken.Token,
		}); err != nil {
			return err
		}
		i++
	}

	if i == 0 {
		return errors.New("rows not found for this query")
	}

	return nil
}
