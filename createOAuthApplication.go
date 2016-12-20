package hitGox

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// CreateOAuthApplication creates a OAuth Application.
func (account *Account) CreateOAuthApplication(name string, redirectURI string) (*Status, error) {
	if err := checkCreateOAuthApplication(account, name, redirectURI); err != nil {
		return nil, err
	}

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

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprint(API, "/oauthapps/", account.UserName)
	resp, err := post(dst, url, &args)
	if err != nil {
		return nil, err
	}

	status, err := fuckYouNeedDecodeStatusFirst(resp)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func checkCreateOAuthApplication(account *Account, name string, redirectURI string) error {
	switch {
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case account.UserName == "":
		return errors.New("username in account can not be empty")
	case name == "":
		return errors.New("name can not be empty")
	case redirectURI == "":
		return errors.New("redirecturi can not be empty")
	}
	return nil
}
