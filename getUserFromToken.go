package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// GetUserFromToken returns user associated with authToken.
func GetUserFromToken(authToken string) (string, error) {
	if authToken == "" {
		return "", errors.New("authtoken can not be empty")
	}

	url := fmt.Sprint(API, "/userfromtoken/", authToken)
	resp, err := get(url, nil)
	if err != nil {
		return "", err
	}

	var obj = struct {
		UserName string `json:"user_name"`
	}{}
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return "", err
	}

	return obj.UserName, nil
}
