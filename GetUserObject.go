package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// User is about basic information about user.
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
//
// When a user isn’t found, this API returns a regular response but with all values containing null.
func GetUserObject(userName, authToken string) (*User, error) {
	var args f.Args
	if authToken != "" {
		args.Add("authToken", authToken)
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("user/", userName))
	body, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj User
	json.NewDecoder(bytes.NewReader(body)).Decode(&obj)

	return &obj, nil
}
