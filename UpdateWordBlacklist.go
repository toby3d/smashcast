package hitGox

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// UpdateWordBlacklist Updates the word blacklist for channel’s chat.
//
// You must send previous blacklist phrases or they will be removed.
func (account *Account) UpdateWordBlacklist(channel string, words []string) (*Status, error) {
	var changes = struct {
		Blacklist []string `json:"blacklist"`
	}{words}

	dst, err := json.Marshal(&changes)
	if err != nil {
		return nil, err
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("chat/blacklist/", channel))
	resp, err := post(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return fixStatus(resp), nil
}
