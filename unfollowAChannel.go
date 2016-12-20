package hitGox

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// UnfollowAChannel removes follower relationship for user id (aka channel).
func (account *Account) UnfollowAChannel(id interface{}) (*Status, error) {
	if err := checkUnfollowAChannel(account, id); err != nil {
		return nil, err
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)
	switch i := id.(type) {
	case int:
		args.Add("follow_id", strconv.Itoa(i))
	case string:
		args.Add("follow_id", i)
	default:
		return nil, errors.New("id can be only as string or int")
	}
	args.Add("type", "user")

	url := fmt.Sprint(API, "/follow")
	resp, err := delete(url, &args)
	if err != nil {
		return nil, err
	}

	status, err := fuckYouNeedDecodeStatusFirst(resp)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func checkUnfollowAChannel(account *Account, id interface{}) error {
	switch {
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case id == nil:
		return errors.New("id can not be empty")
	}
	return nil
}
