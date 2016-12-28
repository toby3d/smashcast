package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// GetTotalViews returns Total Media Views for channel.
func GetTotalViews(channel string) (int, error) {
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("/media/views/", channel))
	resp, err := get(url, nil)
	if err != nil {
		return 0, err
	}

	var obj = struct {
		TotalLiveViews interface{} `json:"total_live_views"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	switch views := obj.TotalLiveViews.(type) {
	case string:
		num, err := strconv.Atoi(views)
		if err != nil {
			return 0, err
		}
		return num, nil
	case bool:
		return 0, nil
	default:
		return 0, errors.New("invalid response format")
	}
}
