package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// GetStreamKey get the stream key for channel.
//
// Editors can read this API.
func (account *Account) GetStreamKey(channel string) (string, error) {
	if err := checkGetStreamKey(account, channel); err != nil {
		return "", err
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprint(API, "/mediakey/", channel)
	resp, err := get(url, &args)
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

func checkGetStreamKey(account *Account, channel string) error {
	switch {
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case channel == "":
		return errors.New("channel can not be empty")
	}
	return nil
}
