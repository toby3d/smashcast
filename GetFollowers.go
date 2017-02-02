package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// Team is about group information.
type Followers struct {
	Request struct {
		This string `json:"this"`
	} `json:"request"`
	Followers []struct {
		Followers      string `json:"followers"`
		UserName       string `json:"user_name"`
		UserID         string `json:"user_id"`
		UserLogo       string `json:"user_logo"`
		UserLogoSmall  string `json:"user_logo_small"`
		FollowID       string `json:"follow_id"`
		FollowerUserID string `json:"follower_user_id"`
		FollowerNotify string `json:"follower_notify"`
		DateAdded      string `json:"date_added"`
	} `json:"followers"`
	MaxResults string `json:"max_results"`
}

// GetFollowers returns a list of users following channel.
func GetFollowers(channel string, offset, limit int, reverse bool, sort []string) (*Followers, error) {
	var args f.Args
	switch {
	case offset > 0:
		args.Add("offset", strconv.Itoa(offset))
		fallthrough
	case limit > 0:
		args.Add("limit", strconv.Itoa(limit))
		fallthrough
	case reverse:
		args.Add("reverse", strconv.FormatBool(reverse))
		fallthrough
	case sort != nil:
		args.Add("sort", fmt.Sprintf("%s", strings.Join(sort, ",")))
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("followers/user/", channel))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj Followers
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
