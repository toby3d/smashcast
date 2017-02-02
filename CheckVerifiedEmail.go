package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
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
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("user/checkVerifiedEmail/", userName))
	resp, err := just.GET(url, nil)
	if err != nil {
		return nil, err
	}

	var obj VerifiedStatus
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
