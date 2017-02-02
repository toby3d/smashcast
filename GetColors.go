package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
)

// GetChatColors get valid chat colors you can use.
func GetChatColors() ([]string, error) {
	url := fmt.Sprintf(APIEndpoint, "chat/colors")
	resp, err := just.GET(url, nil)
	if err != nil {
		return nil, err
	}

	var obj = struct {
		Colors []string `json:"colors"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj.Colors, nil
}
