package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// GetChatBlacklist returns the word blacklist for channel’s chat
func GetChatBlacklist(channel string) (*[]string, error) {
	url := fmt.Sprint(API, "/chat/blacklist/", channel)
	_, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return nil, err
	}

	var obj []string
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
