package hitGox

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// FollowAChannel follows a channel.
//
// id can be either a username or user_id of a user you want to follow.
func (account *Account) FollowAChannel(id interface{}) (*Status, error) {
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
		return nil, errors.New("id can be only as string or int")
	}

	dst, err := json.Marshal(changes)
	if err != nil {
		return nil, err
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, "follow")
	resp, err := post(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return fixStatus(resp), nil
}
