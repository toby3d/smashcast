package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// CreateOAuthApplication creates a OAuth Application.
func (account *Account) CreateOAuthApplication(name string, redirectURI string) (*Status, error) {
	switch {
	case account.AuthToken == "":
		return nil, errors.New("authtoken in account can not be empty")
	case account.UserName == "":
		return nil, errors.New("username in account can not be empty")
	case name == "":
		return nil, errors.New("appname can not be empty")
	case redirectURI == "":
		return nil, errors.New("appredirecturi can not be empty")
	}

	var changes = struct {
		AuthToken string `json:"authToken"`
		UserName  string `json:"user_name"`
		App       struct {
			Name        string `json:"app_name"`
			RedirectURI string `json:"app_redirect_uri"`
		} `json:"app"`
	}{
		AuthToken: account.AuthToken,
		UserName:  account.UserName,
	}

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

	var obj Status
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
