package hitGox

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

// Exchange the exchange request_token for an authToken once the user is redirected.
//
// This access_token is the equivalent of a regular authToken. You are now able to use this on the hitbox API just like any other token.
func (app *Application) Exchange(requestToken string) (string, error) {
	if err := checkExchange(app, requestToken); err != nil {
		return "", err
	}

	hash := base64.StdEncoding.EncodeToString([]byte(app.Token + app.Secret))

	var changes = struct {
		RequestToken string `json:"request_token"`
		AppToken     string `json:"app_token"`
		Hash         string `json:"hash"`
	}{requestToken, app.Token, hash}

	dst, err := json.Marshal(&changes)
	if err != nil {
		return "", err
	}

	url := fmt.Sprint(API, "/oauth/exchange")
	resp, err := post(dst, url, nil)
	if err != nil {
		return "", err
	}

	var obj = struct {
		AccessToken string `json:"access_token"`
	}{}
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return "", err
	}

	return obj.AccessToken, nil
}

func checkExchange(app *Application, requestToken string) error {
	switch {
	case app.Token == "":
		return errors.New("no token of application, create new application first")
	case app.Secret == "":
		return errors.New("no secret of application, create new application first")
	case requestToken == "":
		return errors.New("requesttoken can not be empty")
	}
	return nil
}
