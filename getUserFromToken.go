package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// UserName is a hitbox Username.
type UserName struct {
	UserName string `json:"user_name,ommitempty"`
}

// GetUserFromToken returns user associated with authToken.
func GetUserFromToken(token Token) (UserName, error) {
	requestURL := fmt.Sprintf("%s/userfromtoken/%s", API, token.Token)
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
