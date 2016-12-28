package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// GetStreamKey get the stream key for channel.
//
// Editors can read this API.
func (account *Account) GetStreamKey(channel string) (string, error) {
	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, "/mediakey/"+channel)
	resp, err := get(url, &args)
	if err != nil {
		return "", err
	}

	var obj = struct {
		StreamKey string `json:"streamKey"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj.StreamKey, nil
}
