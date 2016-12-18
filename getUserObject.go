package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// User is about basic information about user.
//
// When a user isn’t found, this API returns a regular response but with all values containing null.
type User struct {
	Followers       string `json:"followers"`
	LivestreamCount string `json:"livestream_count"`
	MediaIsLive     string `json:"media_is_live"`
	MediaLiveSince  string `json:"media_live_since"`
	PartnerType     string `json:"partner_type"`
	Recordings      string `json:"recordings"`
	Teams           string `json:"teams"`
	TFAActive       string `json:"tfa_active"`
	TwitterAccount  string `json:"twitter_account"`
	TwitterEnabled  string `json:"twitter_enabled"`
	UserBetaProfile string `json:"user_beta_profile"`
	UserCover       string `json:"user_cover"`
	UserEmail       string `json:"user_email"`
	UserID          string `json:"user_id"`
	UserLogo        string `json:"user_logo"`
	UserLogoSmall   string `json:"user_logo_small"`
	UserMediaID     string `json:"user_media_id"`
	UserName        string `json:"user_name"`
	UserPartner     string `json:"user_partner"`
	UserStatus      string `json:"user_status"`
	Videos          string `json:"videos"`
}

// GetUserObject returns a regular response about user.
func GetUserObject(userName string, authToken string) (*User, error) {
	var args fasthttp.Args
	args.Add("authToken", authToken)

	url := fmt.Sprintf("%s/user/%s?%s", APIEndpoint, userName, args.String())
	_, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}

	var obj User
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
