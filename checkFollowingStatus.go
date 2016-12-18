package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

// FollowingStatus is about follower relationship.
type FollowingStatus struct {
	Following struct {
		FollowID       string `json:"follow_id"`
		FollowerUserID string `json:"follower_user_id"`
		FollowerNotify string `json:"follower_notify"` // The follower_notify property shows whether that user has email notification turned on 1 or not 0.
	} `json:"following"`
}

// CheckFollowingStatus returns follower relationship from userName to channel.
func CheckFollowingStatus(channel string, userName string) (*FollowingStatus, error) {
	switch {
	case userName == "":
		return nil, errors.New("username can not be empty")
	case channel == "":
		return nil, errors.New("channel can not be empty")
	}

	var args fasthttp.Args
	args.Add("user_name", userName)

	url := APIEndpoint + "/following/user/" + channel
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	var obj FollowingStatus
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
