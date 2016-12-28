package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// ResetStreamKey sets a new stream key for channel.
//
// Editors can run this API.
func (account *Account) ResetStreamKey(channel string) (string, error) {
	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("mediakey/", channel))
	resp, err := put(nil, url, &args)
	if err != nil {
		return "", err
	}

	var obj = struct {
		StreamKey string `json:"streamKey"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj.StreamKey, nil
}
