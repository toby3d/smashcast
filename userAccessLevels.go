package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// Access is a list about user access levels.
type Access struct {
	UserID        string `json:"user_id"`
	AccessUserID  string `json:"access_user_id"`
	Settings      string `json:"settings"`
	Account       string `json:"account"`
	Livestreams   string `json:"livestreams"`
	Partner       string `json:"partner,ommitempty"`
	Broadcast     string `json:"broadcast"`
	Videos        string `json:"videos"`
	Recordings    string `json:"recordings"`
	Statistics    string `json:"statistics"`
	Inbox         string `json:"inbox"`
	Revenues      string `json:"revenues"`
	Chat          string `json:"chat"`
	Following     string `json:"following"`
	Teams         string `json:"teams"`
	Subscriptions string `json:"subscriptions"`
	Admin         string `json:"admin,ommitempty"`
	SuperAdmin    string `json:"superadmin,ommitempty"`
	Payments      string `json:"payments"`
	IsSubscriber  bool   `json:"isSubscriber"`
	IsFollower    bool   `json:"isFollower"`
}

// UserAccessLevels return access levels that authToken has in channel.
func UserAccessLevels(channel string, token Token) (Access, error) {
	requestURL := fmt.Sprintf("%s /user/access/%s/%s", API, channel, token.Token)
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return Access{}, err
	}
	var obj Access
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Access{}, err
	}
	return obj, nil
}
