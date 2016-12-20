package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// ResetStreamKey sets a new stream key for channel.
//
// Editors can run this API.
func (account *Account) ResetStreamKey(channel string) (string, error) {
	if err := checkResetStreamKey(account, channel); err != nil {
		return "", err
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprint(API, "/mediakey/", channel)
	resp, err := put(nil, url, &args)
	if err != nil {
		return "", err
	}

	var obj = struct {
		StreamKey string `json:"streamKey"`
	}{}
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return "", err
	}

	return obj.StreamKey, nil
}

func checkResetStreamKey(account *Account, channel string) error {
	switch {
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case channel == "":
		return errors.New("channel can not be empty")
	}
	return nil
}
