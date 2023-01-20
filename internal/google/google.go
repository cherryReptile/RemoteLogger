package google

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

type User struct {
	ID       string `json:"id" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Verified bool   `json:"verified_email" validate:"required"`
	Picture  string `json:"picture" validate:"required"`
}

func requestToGoogle(token string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)

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

func GetGoogleUserAndBody(token string) (string, []byte, error) {
	user := new(User)
	validate := validator.New()
	res, err := requestToGoogle(token)
	if err != nil {
		return "", nil, err
	}

	if res.StatusCode != http.StatusOK {
		return "", nil, errors.New("google oauth2 failed because returning not ok code")
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

	return user.Email, body, nil
}
