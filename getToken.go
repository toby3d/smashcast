package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
)

// AuthToken is about generated key for most actions by user.
type AuthToken struct {
	AuthToken string `json:"authToken"`
}

// GetToken can be used alternatively to just get an authentication token rather than account information.
func (app *Application) GetToken(login string, pass string) (*string, error) {
	var aToken string

	switch {
	case app.Name == "":
		return &aToken, errors.New("no name of application, create new application first")
	case login == "":
		return &aToken, errors.New("login can not be empty")
	case pass == "":
		return &aToken, errors.New("pass can not be empty")
	}

	var changes = struct {
		Login string `json:"login"`
		Pass  string `json:"pass"`
		App   string `json:"app"`
	}{login, pass, app.Name}

	dst, err := json.Marshal(changes)
	if err != nil {
		return &aToken, err
	}

	url := APIEndpoint + "/auth/token"
	resp, err := post(dst, url, nil)
	if err != nil {
		return &aToken, err
	}

	var obj AuthToken
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return &aToken, err
	}

	aToken = obj.AuthToken

	return &aToken, nil
}
