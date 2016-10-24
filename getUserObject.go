package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

type User struct {
	Followers       string    `json:"followers,ommitempty"`
	IsLive          string    `json:"is_live,ommitempty"`
	LiveSince       time.Time `json:"live_since,ommitempty"`
	LivestreamCount string    `json:"livestream_count,ommitempty"`
	MediaIsLive     string    `json:"media_is_live,ommitempty"`
	MediaLiveSince  time.Time `json:"media_live_since,ommitempty"`
	PartnerType     string    `json:"partner_type,ommitempty"`
	Recordings      string    `json:"recordings,ommitempty"`
	Teams           string    `json:"teams,ommitempty"`
	TwitterAccount  string    `json:"twitter_account,ommitempty"`
	TwitterEnabled  string    `json:"twitter_enabled,ommitempty"`
	BetaProfile     string    `json:"user_beta_profile,ommitempty"`
	Cover           string    `json:"user_cover,ommitempty"`
	Email           string    `json:"user_email,ommitempty"`
	ID              string    `json:"user_id,ommitempty"`
	IsBroadcaster   bool      `json:"user_is_broadcaster"`
	Logo            string    `json:"user_logo,ommitempty"`
	LogoSmall       string    `json:"user_logo_small,ommitempty"`
	MediaID         string    `json:"user_media_id,ommitempty"`
	UserName        string    `json:"user_name,ommitempty"`
	Partner         string    `json:"user_partner,ommitempty"`
	Status          string    `json:"user_status,ommitempty"`
	Videos          string    `json:"videos,ommitempty"`
}

// GetUserObject return information about user.
// When a user isn’t found, this API returns a regular response but with all values containing `null`.
func (token Token) GetUserObject(userName UserName) (User, error) {
	var args fasthttp.Args
	args.Add("authToken", token.Token)
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
