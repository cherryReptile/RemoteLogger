package requests

type GitHubToken struct {
	Token string `json:"token" binding:"required"`
}
