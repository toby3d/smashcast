package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type (
	VerifiedStatus struct {
		Request VerifiedRequest `json:"request"`
		User    VerifiedUser    `json:"user"`
	}

	VerifiedRequest struct {
		This string `json:"this"`
	}

	VerifiedUser struct {
		UserActivated string `json:"user_activated"`
	}
)

// CheckVerifiedEmail check if user has validated their email address.
func CheckVerifiedEmail(userName UserName) (VerifiedStatus, error) {
	requestURL := fmt.Sprintf("%s/user/checkVerifiedEmail/%s", API, userName.UserName)
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return VerifiedStatus{}, err
	}
	var obj VerifiedStatus
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return VerifiedStatus{}, err
	}
	return obj, nil
}
