package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GetUserFromToken returns user associated with authToken.
func GetUserFromToken(authToken string) (string, error) {
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("userfromtoken/", authToken))
	resp, err := get(url, nil)
	if err != nil {
		return "", err
	}

	var obj = struct {
		UserName string `json:"user_name"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj.UserName, nil
}
