package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// UserName is a hitbox Username.
type UserName struct {
	UserName string `json:"user_name,omitempty"`
}

// GetUserFromToken returns user associated with authToken.
func GetUserFromToken(authToken AuthToken) (UserName, error) {
	requestURL := fmt.Sprintf("%s/userfromauthToken/%s", API, authToken.AuthToken)
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return UserName{}, err
	}
	var obj UserName
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return UserName{}, err
	}
	return obj, nil
}
