package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// AuthToken is a Users Authentication Token
type AuthToken struct {
	Token `json:"authToken"`
}

// GetToken get an authentication authToken rather than account information.
func GetToken(login string, pass string, app Application) (AuthToken, error) {
	args := fasthttp.AcquireArgs()
	args.Add("login", login)
	args.Add("pass", pass)
	args.Add("app", app.Name)
	statusCode, body, err := fasthttp.Post(nil, API+"/auth/authToken", args)
	if statusCode != 200 || err != nil {
		return AuthToken{}, err
	}
	var obj AuthToken
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return AuthToken{}, err
	}
	return obj, nil
}
