package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
)

// LastCommercial contains information about last commercial break.
type LastCommercial struct {
	AdCount    string `json:"ad_count"`
	SecondsAgo string `json:"seconds_ago"`
	Timeout    int    `json:"timeout"`
}

// GetLastCommercial returns last commercial break object.
func GetLastCommercial(channel string) (*LastCommercial, error) {
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("ws/combreak/", channel))
	resp, err := just.GET(url, nil)
	if err != nil {
		return nil, err
	}

	var obj LastCommercial
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
