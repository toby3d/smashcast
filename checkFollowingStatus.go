package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type (
	FollowingStatus struct {
		Following Following `json:"following"`
	}

	Following struct {
		FollowID       string `json:"follow_id"`
		FollowerUserID string `json:"follower_user_id"`
		FollowerNotify string `json:"follower_notify"`
	}
)

// CheckFollowingStatus returns follower relationship from `userName` to `channel`
func (userName UserName) CheckFollowingStatus(channel string) (FollowingStatus, error) {
	requestURL := fmt.Sprintf("%s/following/user/%s?user_name=%s", API, channel, userName.UserName)
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
