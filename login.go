package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type (
	Credentials struct {
		Login    string
		Password string
		App      string
		AuthToken
	}

	// Account is a response body about current user account.
	Account struct {
		UserID            string `json:"user_id"`
		UserName          string `json:"user_name"`
		UserLogo          string `json:"user_logo"`
		UserLogoSmall     string `json:"user_logo_small"`
		UserBanned        string `json:"user_banned"`
		UserPartner       string `json:"user_partner,omitempty"`
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
		UserPartner       string `json:"user_partner,omitempty"`
		UserBannedChannel string `json:"user_banned_channel"`
		SuperAdmin        string `json:"superadmin"`
		LivestreamCount   string `json:"livestream_count"`
		Followers         string `json:"followers"`
		AuthToken         string `json:"authToken"`
		Login             string `json:"login"`
		Application       string `json:"app"`
	}
)

// Login used for authentication by user login and password.
func Login(req Credentials) (Account, error) {
	args := fasthttp.AcquireArgs()
	args.Add("login", req.Login)
	args.Add("pass", req.Password)
	args.Add("app", req.App)
	args.Add("authToken", req.AuthToken.Token)
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
