package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// GetToken return authentication token rather than account information.
func (app *Application) GetToken(login string, pass string) (string, error) {
	switch {
	case app.Name == "":
		return "", errors.New("no name of application, create new application first")
	case login == "":
		return "", errors.New("login can not be empty")
	case pass == "":
		return "", errors.New("pass can not be empty")
	}

	var changes = struct {
		Login string `json:"login"`
		Pass  string `json:"pass"`
		App   string `json:"app"`
	}{login, pass, app.Name}

	dst, err := json.Marshal(changes)
	if err != nil {
		return "", err
	}

	url := fmt.Sprint(API, "/auth/token")
	resp, err := post(dst, url, nil)
	if err != nil {
		return "", err
	}

	var obj = struct {
		AuthToken string `json:"authToken"`
	}{}
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return "", err
	}

	return obj.AuthToken, nil
}
