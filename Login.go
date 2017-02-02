package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
)

// Account is about authentication of user account.
type Account struct {
	Access    string `json:"access"`
	App       string `json:"app"`
	AuthToken string `json:"authToken"`
	/*
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
	*/
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
func (app *OAuthApplication) Login(authToken, login, pass string) (*Account, error) {
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

	url := fmt.Sprintf(APIEndpoint, "auth/login")
	resp, err := just.POST(dst, url, nil)
	if err != nil {
		return nil, err
	}

	var obj Account
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
