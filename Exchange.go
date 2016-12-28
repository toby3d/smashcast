package hitGox

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// Exchange the exchange request_token for an authToken once the user is redirected.
//
// This access_token is the equivalent of a regular authToken. You are now able to use this on the hitbox API just like any other token.
func (app *OAuthApplication) Exchange(requestToken string) (string, error) {
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

	url := fmt.Sprintf(APIEndpoint, "oauth/exchange")
	resp, err := post(dst, url, nil)
	if err != nil {
		return "", err
	}

	var obj = struct {
		AccessToken string `json:"access_token"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj.AccessToken, nil
}
