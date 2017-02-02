package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// ResetStreamKey sets a new stream key for channel.
//
// Editors can run this API.
func (account *Account) ResetStreamKey(channel string) (string, error) {
	var args f.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("mediakey/", channel))
	resp, err := just.PUT(nil, url, &args)
	if err != nil {
		return "", err
	}

	var obj = struct {
		StreamKey string `json:"streamKey"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj.StreamKey, nil
}
