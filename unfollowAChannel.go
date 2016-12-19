package hitGox

import (
	"errors"
	"github.com/valyala/fasthttp"
)

// UnfollowAChannel removes follower relationship for user id (aka channel).
func (account *Account) UnfollowAChannel(id string) (*Status, error) {
	switch {
	case account.AuthToken == "":
		return nil, errors.New("authtoken in account can not be empty")
	case id == "":
		return nil, errors.New("id can not be empty")
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)
	args.Add("follow_id", id)
	args.Add("type", "user")

	url := APIEndpoint + "/follow"
	resp, err := delete(url, &args)
	if err != nil {
		return nil, err
	}

	status, err := stupidFuckingStatusResponseByLazyAPIDevelopers(&resp)
	if err != nil {
		return nil, err
	}

	return status, nil
}
