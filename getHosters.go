package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type (
	// HostersList is a response body about hosters.
	HostersList struct {
		Hosters []Hoster `json:"hosters"`
	}

	// Hoster is a single list item about hoster.
	Hoster struct {
		UserName
		UserLogo string `json:"user_logo"`
	}
)

// GetHosters returns channel hosters.
//
// When a user isn’t found, this API returns a regular response but with all values containing null.
//
// Editors can read this API.
func GetHosters(channel string, token Token) (HostersList, error) {
	var args fasthttp.Args
	args.Add("authToken", token.Token)
	requestURL := fmt.Sprintf("%s/hosters/%s?%s", API, channel, args.String())
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return HostersList{}, err
	}
	var obj HostersList
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return HostersList{}, err
	}
	return obj, nil
}
