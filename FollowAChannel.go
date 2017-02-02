package hitGox

import (
	"encoding/json"
	"fmt"
	"strconv"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// FollowAChannel follows a channel.
//
// id can be either a username or user_id of a user you want to follow.
func (account *Account) FollowAChannel(id interface{}) (*just.Status, error) {
	var changes = struct {
		Type     string `json:"type"`
		FollowID string `json:"follow_id"`
	}{Type: "user"}

	switch i := id.(type) {
	case int:
		changes.FollowID = strconv.Itoa(i)
	case string:
		changes.FollowID = i
	default:
		return nil, fmt.Errorf("id can be only as string or int")
	}

	dst, err := json.Marshal(changes)
	if err != nil {
		return nil, err
	}

	var args f.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, "follow")
	resp, err := just.POST(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
