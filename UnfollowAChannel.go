package hitGox

import (
	"fmt"
	"strconv"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// UnfollowAChannel removes follower relationship for user id (aka channel).
func (account *Account) UnfollowAChannel(id interface{}) (*just.Status, error) {
	var args f.Args
	args.Add("authToken", account.AuthToken)
	switch i := id.(type) {
	case int:
		args.Add("follow_id", strconv.Itoa(i))
	case string:
		args.Add("follow_id", i)
	default:
		return nil, fmt.Errorf("id can be only as string or int")
	}
	args.Add("type", "user")

	url := fmt.Sprintf(APIEndpoint, "follow")
	resp, err := just.DELETE(url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
