package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// ListOAuthAccess contains information about connectet AOuth Applications on channel.
type ListOAuthAccess struct {
	Apps []struct {
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"apps"`
}

// GetListOAuthAccess returns OAuth Applications the user has authenticated with.
func (account *Account) GetListOAuthAccess() (*ListOAuthAccess, error) {
	if err := checkGetListOAuthAccess(account); err != nil {
		return nil, err
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprint(API, "/oauthaccess/", account.UserName)
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	var obj ListOAuthAccess
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}

func checkGetListOAuthAccess(account *Account) error {
	switch {
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case account.UserName == "":
		return errors.New("username in account can not be empty")
	}
	return nil
}
