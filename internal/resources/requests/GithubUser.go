package requests

type GithubUser struct {
	ID        uint   `json:"id" binding:"required"`
	Login     string `json:"login" binding:"required"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url" binding:"required"`
}
