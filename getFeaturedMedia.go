package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// FeaturedMedia is about featured live stream.
type FeaturedMedia struct {
	MediaID          string `json:"media_id"`
	MediaDisplayName string `json:"media_display_name"`
	MediaName        string `json:"media_name"`
	Backdrop         string `json:"backdrop"`
	BackdropHTML     string `json:"backdrop_html"`
}

// GetFeaturedMedia returns a featured live stream.
func GetFeaturedMedia() (*FeaturedMedia, error) {
	url := fmt.Sprint(API, "/mediafeatured")
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj FeaturedMedia
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
