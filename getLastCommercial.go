package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// LastCommercial contains information about last commercial break.
type LastCommercial struct {
	AdCount    string `json:"ad_count"`
	SecondsAgo string `json:"seconds_ago"`
	Timeout    int    `json:"timeout"`
}

// GetLastCommercial returns last commercial break object.
func GetLastCommercial(channel string) (*LastCommercial, error) {
	if channel == "" {
		return nil, errors.New("channel can not be empty")
	}

	url := fmt.Sprint(API, "/ws/combreak/", channel)
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj LastCommercial
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
