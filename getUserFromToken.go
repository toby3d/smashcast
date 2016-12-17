package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// UserName is about user username.
type UserName struct {
	UserName string `json:"user_name"`
}

// GetUserFromToken returns user associated with authToken.
func GetUserFromToken(authToken string) (*UserName, error) {
	if authToken == "" {
		return nil, errors.New("this action requires a authtoken")
	}

	url := fmt.Sprintf("%s/userfromtoken/%s", APIEndpoint, authToken)
	_, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}

	var obj UserName
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
