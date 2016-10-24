package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type (
	HostersList struct {
		Hosters []Hosters `json:"hosters"`
	}

	Hosters struct {
		UserName string `json:"user_name"`
		UserLogo string `json:"user_logo"`
	}
)

// GetHosters returns channel hosters.
// When a user isn’t found, this API returns a regular response but with all values containing `null`.
func (token Token) GetHosters(channel string) (HostersList, error) {
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
