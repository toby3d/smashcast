package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// GetChatColors get valid chat colors you can use.
func GetChatColors() ([]string, error) {
	url := fmt.Sprintf(APIEndpoint, "chat/colors")
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj = struct {
		Colors []string `json:"colors"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj.Colors, nil
}
