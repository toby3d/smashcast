package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
)

// UserName is about user username.
type UserName struct {
	UserName string `json:"user_name"`
}

// GetUserFromToken returns user associated with authToken.
func GetUserFromToken(authToken string) (*string, error) {
	var uName string

	if authToken == "" {
		return &uName, errors.New("this action requires a authtoken")
	}

	url := APIEndpoint + "/userfromtoken/" + authToken
	resp, err := get(url, nil)
	if err != nil {
		return &uName, err
	}

	var obj UserName
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return &uName, err
	}

	uName = obj.UserName

	return &uName, nil
}
