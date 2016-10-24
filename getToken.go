package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type AuthToken struct {
	Token `json:"authToken"`
}

// GetToken get an authentication token rather than account information.
func (app Application) GetToken(login string, pass string) (AuthToken, error) {
	args := fasthttp.AcquireArgs()
	args.Add("login", login)
	args.Add("pass", pass)
	args.Add("app", app.Name)
	statusCode, body, err := fasthttp.Post(nil, API+"/auth/token", args)
	if statusCode != 200 || err != nil {
		return AuthToken{}, err
	}
	var obj AuthToken
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return AuthToken{}, err
	}
	return obj, nil
}
