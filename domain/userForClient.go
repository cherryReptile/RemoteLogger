package domain

type ClientUser struct {
	User
	Profile
	AuthToken
}

type ClientUserRepo interface {
	GetUserWithProfile(clientUser *ClientUser, userID string) error
	GetAuthClientUser(clientUser *ClientUser, userID, token string) error
}

type ClientUserUsecase interface {
	GetUserWithProfile(userAndProfile *ClientUser, userID string) error
	GetAuthClientUser(clientUser *ClientUser, userID, token string) error
}
