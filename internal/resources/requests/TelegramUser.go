package requests

import "time"

type TelegramUser struct {
	ID        uint      `json:"id" binding:"required"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	PhotoURL  string    `json:"photo_url" binding:"required"`
	AuthDate  time.Time `json:"auth_date" binding:"required"`
	Hash      string    `json:"hash" binding:"required"`
}
