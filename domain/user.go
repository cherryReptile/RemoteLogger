package domain

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

//type ClientUser struct {
//	User
//	Profile
//}

type UserRepo interface {
	Create(user *User) error
	Find(user *User, uuid string) error
	FindByLoginAndProvider(user *User, username, provider string) error
	GetTokenByStr(user *User, tokenStr string) (*AuthToken, error)
	//GetUserWithProfile(userAndProfile *ClientUser, userID string) error
}

type UserUsecase interface {
	Create(user *User) error
	Find(user *User, uuid string) error
	FindByLoginAndProvider(user *User, username, provider string) error
	GetTokenByStr(user *User, tokenStr string) (*AuthToken, error)
	//GetUserWithProfile(userAndProfile *ClientUser, userID string) error
}
