package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// User contains information about user.
type User struct {
	Followers       string    `json:"followers,omitempty"`
	IsLive          string    `json:"is_live,omitempty"`
	LiveSince       Timestamp `json:"live_since,omitempty"`
	LivestreamCount string    `json:"livestream_count,omitempty"`
	MediaIsLive     string    `json:"media_is_live,omitempty"`
	MediaLiveSince  Timestamp `json:"media_live_since,omitempty"`
	PartnerType     string    `json:"partner_type,omitempty"`
	Recordings      string    `json:"recordings,omitempty"`
	Teams           string    `json:"teams,omitempty"`
	TwitterAccount  string    `json:"twitter_account,omitempty"`
	TwitterEnabled  string    `json:"twitter_enabled,omitempty"`
	BetaProfile     string    `json:"user_beta_profile,omitempty"`
	Cover           string    `json:"user_cover,omitempty"`
	Email           string    `json:"user_email,omitempty"`
	ID              string    `json:"user_id,omitempty"`
	IsBroadcaster   bool      `json:"user_is_broadcaster"`
	Logo            string    `json:"user_logo,omitempty"`
	LogoSmall       string    `json:"user_logo_small,omitempty"`
	MediaID         string    `json:"user_media_id,omitempty"`
	UserName        string    `json:"user_name,omitempty"`
	Partner         string    `json:"user_partner,omitempty"`
	Status          string    `json:"user_status,omitempty"`
	Videos          string    `json:"videos,omitempty"`
}

// GetUserObject return information about user.
//
// When a user isn’t found, this API returns a regular response but with all values containing null.
func GetUserObject(userName UserName, authToken AuthToken) (User, error) {
	var args fasthttp.Args

	// Returns private user details.
	args.Add("authToken", authToken.Token)

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
