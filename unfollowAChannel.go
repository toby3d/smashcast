package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

	url := fmt.Sprint(API, "/follow")
	resp, err := delete(url, &args)
	if err != nil {
		return nil, err
	}

	var obj Status
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
