package hitGox

import (
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// CheckFollowingjust.Status returns follower relationship from userName to channel.
func (account *Account) UpdateWhisperSetting(whisper bool) (*just.Status, error) {
	var args f.Args
	args.Add("authToken", account.AuthToken)

	var allow string
	if !whisper {
		allow = "0"
	} else {
		allow = "1"
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("chat/pm/", allow))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
