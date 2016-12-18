package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
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

	var changes = struct {
		Login     string `json:"login"`
		Password  string `json:"pass"`
		App       string `json:"app"`
		AuthToken string `json:"authToken"`
	}{login, pass, app.Name, authToken}

	dst, err := json.Marshal(&changes)
	if err != nil {
		return nil, err
	}

	url := APIEndpoint + "/auth/login"
	resp, err := post(dst, url, nil)
	if err != nil {
		return nil, err
	}

	var obj Account
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
