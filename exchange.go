package hitGox

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

// AccessToken is the equivalent of a regular authToken. You are now able to use this on the hitbox API just like any other token.
type AccessToken struct {
	AccessToken string `json:"access_token"`
}

// Exchange the exchange request_token for an authToken once the user is redirected.
func (app *Application) Exchange(requestToken string) (*AccessToken, error) {
	switch {
	case app.Token == "":
		return nil, errors.New("no token of application, create new application first")
	case app.Secret == "":
		return nil, errors.New("no secret of application, create new application first")
	case requestToken == "":
		return nil, errors.New("requesttoken can not be empty")
	}

	hash := base64.StdEncoding.EncodeToString([]byte(app.Token + app.Secret))

	var args fasthttp.Args
	args.Add("request_token", requestToken)
	args.Add("app_token", app.Token)
	args.Add("hash", hash)

	url := APIEndpoint + "/oauth/exchange"
	_, resp, err := fasthttp.Post(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var obj AccessToken
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
