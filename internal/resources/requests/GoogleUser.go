package requests

type GoogleUser struct {
	Email   string `json:"email" binding:"required,email"`
	Picture string `json:"picture" binding:"required"`
}
