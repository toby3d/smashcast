package hitGox

import (
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// UpdateChatSettings update chat settings for channel.
//
// Editors can modify this API, except the whisper setting.
func (account *Account) UpdateChatSettings(channel string, subImages, whisper bool) (*just.Status, error) {
	var changes = struct {
		UserID    string `json:"user_id"`
		SubImages bool   `json:"sub_images"`
		Whisper   bool   `json:"whisper"`
	}{account.UserID, subImages, whisper}

	dst, err := json.Marshal(&changes)
	if err != nil {
		return nil, err
	}

	var args f.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("chat/settings/", channel))
	resp, err := just.POST(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
