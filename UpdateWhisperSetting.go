package hitGox

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// CheckFollowingStatus returns follower relationship from userName to channel.
func (account *Account) UpdateWhisperSetting(whisper bool) (*Status, error) {
	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	var allow string
	if !whisper {
		allow = "0"
	} else {
		allow = "1"
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("chat/pm/", allow))
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	return fixStatus(resp), nil
}
