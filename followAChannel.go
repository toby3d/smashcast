package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

// FollowAChannel follows a channel.
//
// id can be either a username or user_id of a user you want to follow.
func (account *Account) FollowAChannel(id string) (*Status, error) {
	switch {
	case account.AuthToken == "":
		return nil, errors.New("authtoken in account can not be empty")
	case id == "":
		return nil, errors.New("id can not be empty")
	}

	var changes = struct {
		Type     string `json:"type"`
		FollowID string `json:"follow_id"`
	}{"user", id}

	dst, err := json.Marshal(changes)
	if err != nil {
		return nil, err
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := APIEndpoint + "/follow"
	resp, err := post(dst, url, &args)
	if err != nil {
		return nil, err
	}

	var obj Status
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
