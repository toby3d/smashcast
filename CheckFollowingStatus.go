package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
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
	var args f.Args
	args.Add("user_name", userName)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("following/user/", channel))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj FollowingStatus
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
