package hitGox

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
)

// AccessToken is the equivalent of a regular authToken. You are now able to use this on the hitbox API just like any other token.
type AccessToken struct {
	AccessToken string `json:"access_token"`
}

// Exchange the exchange request_token for an authToken once the user is redirected.
func (app *Application) Exchange(requestToken string) (*string, error) {
	var aToken string

	switch {
	case app.Token == "":
		return &aToken, errors.New("no token of application, create new application first")
	case app.Secret == "":
		return &aToken, errors.New("no secret of application, create new application first")
	case requestToken == "":
		return &aToken, errors.New("requesttoken can not be empty")
	}

	hash := base64.StdEncoding.EncodeToString([]byte(app.Token + app.Secret))

	var changes = struct {
		RequestToken string `json:"request_token"`
		AppToken     string `json:"app_token"`
		Hash         string `json:"hash"`
	}{requestToken, app.Token, hash}

	dst, err := json.Marshal(&changes)
	if err != nil {
		return nil, err
	}

	url := APIEndpoint + "/oauth/exchange"
	resp, err := post(dst, url, nil)
	if err != nil {
		return &aToken, err
	}

	var obj AccessToken
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return &aToken, err
	}
	aToken = obj.AccessToken

	return &aToken, nil
}
