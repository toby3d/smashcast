package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// HostersList contains information about hosted channels.
type HostersList struct {
	Hosters []struct {
		UserLogo string `json:"user_logo"`
		UserName string `json:"user_name"`
	} `json:"hosters"`
}

// GetHosters returns a list of channels hosting channel.
//
// Editors can read this API.
func (account *Account) GetHosters(channel string) (*HostersList, error) {
	var args f.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("hosters/", channel))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj HostersList
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
