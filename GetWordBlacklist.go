package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
)

// GetWordBlacklist returns the word blacklist for channel’s chat.
func GetWordBlacklist(channel string) ([]string, error) {
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("chat/blacklist/", channel))
	resp, err := just.GET(url, nil)
	if err != nil {
		return nil, err
	}

	var obj []string
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj, nil
}
