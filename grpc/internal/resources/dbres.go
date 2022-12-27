package resources

type AppUserData struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}
