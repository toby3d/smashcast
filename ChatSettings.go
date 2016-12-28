package hitGox

import (
	"bytes"
	"encoding/json"
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
	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("chat/settings/", channel))
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	var obj ChatSettings
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
