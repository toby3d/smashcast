package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type (
	// FollowingStatus is a response body.
	FollowingStatus struct {
		Following Follower `json:"following"`
	}

	// Follower show information about following status.
	Follower struct {
		FollowID       string `json:"follow_id"`
		FollowerUserID string `json:"follower_user_id"`

		// The follower_notify property shows whether that user has email notification turned on 1 or not 0.
		FollowerNotify string `json:"follower_notify"`
	}
)

// CheckFollowingStatus returns follower relationship from userName to channel.
func CheckFollowingStatus(userName UserName, channel string) (FollowingStatus, error) {
	var args fasthttp.Args
	args.Add("user_name", userName.UserName)
	requestURL := fmt.Sprintf("%s/following/user/%s?%s", API, channel, args.String())
	_, resp, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return FollowingStatus{}, err
	}
	var obj FollowingStatus
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return FollowingStatus{}, err
	}
	return obj, nil
}
