package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
)

// Livejust.Status contains info about live status of channel.
type LiveStatus struct {
	MediaIsLive string `json:"media_is_live"`
	MediaViews  string `json:"media_views"`
}

// GetLivejust.Status Returns media status and viewer count for channel.
func GetLiveStatus(channel string) (*LiveStatus, error) {
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("media/status/", channel))
	resp, err := just.GET(url, nil)
	if err != nil {
		return nil, err
	}

	var obj LiveStatus
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
