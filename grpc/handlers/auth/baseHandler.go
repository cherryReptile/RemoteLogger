package auth

import (
	"database/sql"
	"encoding/json"
	"github.com/cherryReptile/WS-AUTH/domain"
)

type BaseHandler struct {
	userUsecase           domain.UserUsecase
	providerUsecase       domain.ProviderUsecase
	tokenUsecase          domain.AuthTokenUsecase
	providersDataUsecase  domain.ProvidersDataUsecase
	usersProvidersUsecase domain.UsersProvidersUsecase
	profileUsecase        domain.ProfileUsecase
}

func (h *BaseHandler) SetProfile(profile *domain.Profile, userID string) error {
	profile.FirstName = sql.NullString{Valid: true, String: ""}
	profile.LastName = sql.NullString{Valid: true, String: ""}
	profile.Address = sql.NullString{Valid: true, String: ""}
	profile.UserID = userID
	od, err := json.Marshal(map[string]string{"": ""})
	if err != nil {
		return err
	}
	profile.OtherData = od
	return nil
}
