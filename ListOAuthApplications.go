package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

type (
	ListOAuthApplications struct {
		Success bool               `json:"success"`
		Apps    []OAuthApplication `json:"apps"`
	}

	OAuthApplication struct {
		ID          string `json:"app_id"`
		Name        string `json:"app_name"`
		Note        string `json:"app_note"`
		Enabled     string `json:"app_enabled"`
		RedirectURI string `json:"app_redirect_uri"`
		Token       string `json:"app_token"`
		Secret      string `json:"app_secret"`
		UserID      string `json:"app_user_id"`
	}
)

func (acc *Account) GetListOAuthApplications() (*ListOAuthApplications, error) {
	var args f.Args
	args.Add("authToken", acc.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("oauthapps/", acc.UserName))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj ListOAuthApplications
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
