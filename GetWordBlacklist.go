package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GetWordBlacklist returns the word blacklist for channel’s chat.
func GetWordBlacklist(channel string) ([]string, error) {
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("chat/blacklist/", channel))
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj []string
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj, nil
}
