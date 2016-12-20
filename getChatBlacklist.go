package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// GetChatBlacklist returns the word blacklist for channel’s chat
func GetChatBlacklist(channel string) ([]string, error) {
	if channel == "" {
		return nil, errors.New("channel can not be empty")
	}

	url := fmt.Sprint(API, "/chat/blacklist/", channel)
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj []string
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return obj, nil
}
