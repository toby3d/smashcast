package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type (
	// Account is a response body about current user account.
	Account struct {
		UserID            string `json:"user_id"`
		UserName          string `json:"user_name"`
		UserLogo          string `json:"user_logo"`
		UserLogoSmall     string `json:"user_logo_small"`
		UserBanned        string `json:"user_banned"`
		UserPartner       string `json:"user_partner,ommitempty"`
		UserBannedChannel string `json:"user_banned_channel"`
		SuperAdmin        string `json:"superadmin"`
		LivestreamCount   string `json:"livestream_count"`
		Followers         string `json:"followers"`
		AuthToken         string `json:"authToken"`
		Login             string `json:"login"`
		Data              Data   `json:"data"`
		Access            string `json:"access"`
		Application       string `json:"app"`
	}

	// Data is a part Account response body about current user account.
	Data struct {
		UserID            string `json:"user_id"`
		UserName          string `json:"user_name"`
		UserLogo          string `json:"user_logo"`
		UserLogoSmall     string `json:"user_logo_small"`
		UserBanned        string `json:"user_banned"`
		UserPartner       string `json:"user_partner,ommitempty"`
		UserBannedChannel string `json:"user_banned_channel"`
		SuperAdmin        string `json:"superadmin"`
		LivestreamCount   string `json:"livestream_count"`
		Followers         string `json:"followers"`
		AuthToken         string `json:"authToken"`
		Login             string `json:"login"`
		Application       string `json:"app"`
	}
)

// LoginByCredentials used for authentication by user login and password.
func LoginByCredentials(login UserName, pass string, app Application) (Account, error) {
	args := fasthttp.AcquireArgs()
	args.Add("login", login.UserName)
	args.Add("pass", pass)
	args.Add("app", app.Name)
	statusCode, body, err := fasthttp.Post(nil, API+"/auth/login", args)
	if statusCode != 200 || err != nil {
		return Account{}, err
	}
	var obj Account
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Account{}, err
	}
	return obj, nil
}

// LoginByToken used for authentication by user token.
func LoginByToken(token Token, app Application) (Account, error) {
	args := fasthttp.AcquireArgs()
	args.Add("app", app.Name)
	args.Add("authToken", token.Token)
	statusCode, body, err := fasthttp.Post(nil, API+"/auth/login", args)
	if statusCode != 200 || err != nil {
		return Account{}, err
	}
	var obj Account
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Account{}, err
	}
	return obj, nil
}
