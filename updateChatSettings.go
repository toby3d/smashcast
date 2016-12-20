package hitGox

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// UpdateChatSettings update chat settings for channel.
//
// Editors can modify this API, except the whisper setting.
func (account *Account) UpdateChatSettings(channel string, subImages bool, whisper bool) (*Status, error) {
	if err := checkUpdateChatSettings(account, channel); err != nil {
		return nil, err
	}

	var changes = struct {
		UserID    string `json:"user_id"`
		SubImages bool   `json:"sub_images"`
		Whisper   bool   `json:"whisper"`
	}{account.UserID, subImages, whisper}

	dst, err := json.Marshal(&changes)
	if err != nil {
		return nil, err
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprint(API, "/chat/settings/", channel)
	resp, err := post(dst, url, &args)
	if err != nil {
		return nil, err
	}

	status, err := fuckYouNeedDecodeStatusFirst(resp)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func checkUpdateChatSettings(account *Account, channel string) error {
	switch {
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case channel == "":
		return errors.New("channel can not be empty")
	}
	return nil
}
