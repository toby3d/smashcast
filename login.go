package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

// Account is about authentication of user account.
type Account struct {
	Access    string `json:"access"`
	App       string `json:"app"`
	AuthToken string `json:"authToken"`
	Data      struct {
		App               string `json:"app"`
		AuthToken         string `json:"authToken"`
		Followers         string `json:"followers"`
		LivestreamCount   string `json:"livestream_count"`
		Login             string `json:"login"`
		Superadmin        string `json:"superadmin"`
		UserBanned        string `json:"user_banned"`
		UserBannedChannel string `json:"user_banned_channel"`
		UserID            string `json:"user_id"`
		UserLogo          string `json:"user_logo"`
		UserLogoSmall     string `json:"user_logo_small"`
		UserName          string `json:"user_name"`
		UserPartner       string `json:"user_partner"`
	} `json:"data"`
	Followers         string `json:"followers"`
	LivestreamCount   string `json:"livestream_count"`
	Login             string `json:"login"`
	Superadmin        string `json:"superadmin"`
	UserBanned        string `json:"user_banned"`
	UserBannedChannel string `json:"user_banned_channel"`
	UserID            string `json:"user_id"`
	UserLogo          string `json:"user_logo"`
	UserLogoSmall     string `json:"user_logo_small"`
	UserName          string `json:"user_name"`
	UserPartner       string `json:"user_partner"`
}

// Login authenticates and returns account information.
func (app *Application) Login(login string, pass string, authToken string) (*Account, error) {
	switch {
	case app.Name == "":
		return nil, errors.New("no name of application, create new application first")
	case login == "" && pass == "" && authToken == "":
		return nil, errors.New("empty details, use authtoken or login/pass")
	case (login == "" || pass == "") && authToken == "":
		return nil, errors.New("account details can not be empty")
	}

	var args fasthttp.Args
	args.Add("app", app.Name)
	args.Add("authToken", authToken)
	args.Add("login", login)
	args.Add("pass", pass)

	url := APIEndpoint + "/auth/login"
	_, resp, err := fasthttp.Post(nil, url, &args)
	if err != nil {
		return nil, err
	}

	var obj Account
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
