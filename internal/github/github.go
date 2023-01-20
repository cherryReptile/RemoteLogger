package github

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

type User struct {
	ID        uint   `json:"id" validate:"required"`
	Login     string `json:"login" validate:"required"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url" validate:"required"`
}

func requestToGithub(token string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetGitHubUserAndBody(token string) (string, []byte, error) {
	user := new(User)
	validate := validator.New()
	res, err := requestToGithub(token)
	if err != nil {
		return "", nil, err
	}

	if res.StatusCode != http.StatusOK {
		return "", nil, errors.New("github oauth2 failed because returning not ok code")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", nil, err
	}

	err = json.Unmarshal(body, user)

	if err != nil {
		return "", nil, err
	}

	if err = validate.Struct(user); err != nil {
		return "", nil, err
	}

	return user.Login, body, nil
}
