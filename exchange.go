package hitGox

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type AccessToken struct {
	Token `json:"acces_token"`
}

// Exchange get an authentication token rather than account information.
func (app Application) Exchange(requestToken string) (AccessToken, error) {
	hash := base64.StdEncoding.EncodeToString([]byte(app.Token + app.Secret))

	args := fasthttp.AcquireArgs()
	args.Add("request_token", requestToken)
	args.Add("app_token", app.Token)
	args.Add("hash", hash)
	statusCode, body, err := fasthttp.Post(nil, API+"/oauth/exchange", args)
	if statusCode != 200 || err != nil {
		return AccessToken{}, err
	}
	var obj AccessToken
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return AccessToken{}, err
	}
	return obj, nil
}
