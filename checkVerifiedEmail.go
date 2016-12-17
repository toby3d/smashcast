package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// VerifiedStatus is about validated user email address.
//
// If user object is empty or user_activated property is 0, the user hasn’t verified their email address.
type VerifiedStatus struct {
	Request struct {
		This string `json:"this"`
	} `json:"request"`
	User struct {
		UserActivated string `json:"user_activated"`
	} `json:"user"`
}

// CheckVerifiedEmail check if user has validated their email address.
func CheckVerifiedEmail(userName string) (*VerifiedStatus, error) {
	if userName == "" {
		return nil, errors.New("username can not be empty")
	}

	url := fmt.Sprintf("%s/user/checkVerifiedEmail/%s", APIEndpoint, userName)
	_, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}

	var obj VerifiedStatus
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
