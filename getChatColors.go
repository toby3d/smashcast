package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ChatColors is a list with valid chat HEX-colors.
type ChatColors struct {
	Colors []string `json:"colors"`
}

// GetChatColors get valid chat colors you can use.
func GetChatColors() (*ChatColors, error) {
	url := fmt.Sprint(API, "/chat/colors")
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj ChatColors
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
