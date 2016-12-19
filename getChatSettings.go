package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// ChatSettings containing information about channel chat.
type ChatSettings struct {
	UserID     string   `json:"user_id"`
	SubImages  bool     `json:"sub_images"`
	Whisper    bool     `json:"whisper"`
	IgnoreList []string `json:"ignore_list"`
}

// GetChatSettings returns chat settings for channel.
//
// Moderators and Editors can view this API.
func (account *Account) GetChatSettings(channel string) (*ChatSettings, error) {
	switch {
	case account.AuthToken == "":
		return nil, errors.New("authtoken in account can not be empty")
	case channel == "":
		return nil, errors.New("channel can not be empty")
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprint(API, "/chat/settings/", channel)
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	var obj ChatSettings
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
