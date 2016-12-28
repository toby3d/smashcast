package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// LiveStatus contains info about live status of channel.
type LiveStatus struct {
	MediaIsLive string `json:"media_is_live"`
	MediaViews  string `json:"media_views"`
}

// GetLiveStatus Returns media status and viewer count for channel.
func GetLiveStatus(channel string) (*LiveStatus, error) {
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("media/status/", channel))
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj LiveStatus
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
