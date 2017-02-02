package hitGox

import (
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// UpdateWordBlacklist Updates the word blacklist for channel’s chat.
//
// You must send previous blacklist phrases or they will be removed.
func (account *Account) UpdateWordBlacklist(channel string, words []string) (*just.Status, error) {
	var changes = struct {
		Blacklist []string `json:"blacklist"`
	}{words}

	dst, err := json.Marshal(&changes)
	if err != nil {
		return nil, err
	}

	var args f.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("chat/blacklist/", channel))
	resp, err := just.POST(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
