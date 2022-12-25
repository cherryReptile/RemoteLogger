package requests

type GoogleUser struct {
	ID       string `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Verified bool   `json:"verified_email" binding:"required"`
	Picture  string `json:"picture" binding:"required"`
}
