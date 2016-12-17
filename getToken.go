package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

// AuthToken is about generated key for most actions by user.
type AuthToken struct {
	AuthToken string `json:"authToken"`
}

// GetToken can be used alternatively to just get an authentication token rather than account information.
func (app *Application) GetToken(login string, pass string) (*AuthToken, error) {
	switch {
	case app.Name == "":
		return nil, errors.New("no name of application, create new application first")
	case login == "":
		return nil, errors.New("login can not be empty")
	case pass == "":
		return nil, errors.New("pass can not be empty")
	}

	var args fasthttp.Args
	args.Add("login", login)
	args.Add("pass", pass)
	args.Add("app", app.Name)

	url := APIEndpoint + "/auth/token"
	_, resp, err := fasthttp.Post(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var obj AuthToken
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
