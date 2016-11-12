package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type (
	// Account is a response body about current user account.
	AccountInformation struct {
		UserID            string `json:"user_id,omitempty"`
		UserName          string `json:"user_name,omitempty"`
		UserLogo          string `json:"user_logo,omitempty"`
		UserLogoSmall     string `json:"user_logo_small,omitempty"`
		UserBanned        string `json:"user_banned,omitempty"`
		UserPartner       string `json:"user_partner,omitempty"`
		UserBannedChannel string `json:"user_banned_channel,omitempty"`
		SuperAdmin        string `json:"superadmin,omitempty"`
		LivestreamCount   string `json:"livestream_count,omitempty"`
		Followers         string `json:"followers,omitempty"`
		AuthToken         string `json:"authToken,omitempty"`
		Login             string `json:"login,omitempty"`
		Data              Data   `json:"data"`
		Access            string `json:"access,omitempty"`
		App               string `json:"app,omitempty"`
	}

	// Data is a part Account response body about current user account.
	Data struct {
		UserID            string `json:"user_id,omitempty"`
		UserName          string `json:"user_name,omitempty"`
		UserLogo          string `json:"user_logo,omitempty"`
		UserLogoSmall     string `json:"user_logo_small,omitempty"`
		UserBanned        string `json:"user_banned,omitempty"`
		UserPartner       string `json:"user_partner,omitempty"`
		UserBannedChannel string `json:"user_banned_channel,omitempty"`
		SuperAdmin        string `json:"superadmin,omitempty"`
		LivestreamCount   string `json:"livestream_count,omitempty"`
		Followers         string `json:"followers,omitempty"`
		AuthToken         string `json:"authToken,omitempty"`
		Login             string `json:"login,omitempty"`
		App               string `json:"app,omitempty"`
	}
)

// Login used for authentication by user login and password.
func (authToken AuthToken) Login(app Application) (AccountInformation, error) {
	args := fasthttp.AcquireArgs()
	args.Add("app", app.Name)
	args.Add("authToken", authToken.AuthToken)
	statusCode, body, err := fasthttp.Post(nil, API+"/auth/login", args)
	if statusCode != 200 || err != nil {
		return AccountInformation{}, err
	}
	var obj AccountInformation
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return AccountInformation{}, err
	}
	return obj, nil
}
