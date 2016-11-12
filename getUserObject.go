package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// User contains information about user.
type User struct {
	Followers         string `json:"followers,omitempty"`
	Videos            string `json:"videos,omitempty"`
	Recordings        string `json:"recordings,omitempty"`
	Teams             string `json:"teams,omitempty"`
	UserID            string `json:"user_id,omitempty"`
	UserName          string `json:"user_name,omitempty"`
	UserStatus        string `json:"user_status,omitempty"`
	UserLogo          string `json:"user_logo,omitempty"`
	UserCover         string `json:"user_cover,omitempty"`
	UserLogoSmall     string `json:"user_logo_small,omitempty"`
	UserIsBroadcaster bool   `json:"user_is_broadcaster"`
	UserEmail         string `json:"user_email,omitempty"`
	UserPartner       string `json:"user_partner,omitempty"`
	PartnerType       string `json:"partner_type,omitempty"`
	UserBetaProfile   string `json:"user_beta_profile,omitempty"`
	MediaIsLive       string `json:"media_is_live,omitempty"`
	MediaLiveSince    string `json:"media_live_since,omitempty"`
	UserMediaID       string `json:"user_media_id,omitempty"`
	TwitterAccount    string `json:"twitter_account,omitempty"`
	TwitterEnabled    string `json:"twitter_enabled,omitempty"`
	LivestreamCount   string `json:"livestream_count,omitempty"`
	TFAActive         string `json:"tfa_active,omitempty"`
	IsLive            string `json:"is_live,omitempty"`
	LiveSince         string `json:"live_since,omitempty"`
}

// GetUserObject return information about user.
//
// When a user isn’t found, this API returns a regular response but with all values containing null.
func GetUserObject(userName UserName, authToken AuthToken) (User, error) {
	var args fasthttp.Args

	// Returns private user details.
	args.Add("authToken", authToken.AuthToken)

	requestURL := fmt.Sprintf("%s/user/%s?%s", API, userName.UserName, args.String())
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return User{}, err
	}
	var obj User
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return User{}, err
	}
	return obj, nil
}
