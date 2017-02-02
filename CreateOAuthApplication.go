package hitGox

import (
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// CreateOAuthApplication creates a OAuth Application.
func (account *Account) CreateOAuthApplication(name string, redirectURI string) (*just.Status, error) {
	var changes = struct {
		AuthToken string `json:"authToken"`
		UserName  string `json:"user_name"`
		App       struct {
			Name        string `json:"app_name"`
			RedirectURI string `json:"app_redirect_uri"`
		} `json:"app"`
	}{}

	changes.AuthToken = account.AuthToken
	changes.UserName = account.UserName
	changes.App.Name = name
	changes.App.RedirectURI = redirectURI

	dst, err := json.Marshal(&changes)
	if err != nil {
		return nil, err
	}

	var args f.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("oauthapps/", account.UserName))
	resp, err := just.POST(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
