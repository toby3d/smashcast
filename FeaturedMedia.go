package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
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
	url := fmt.Sprintf(APIEndpoint, "mediafeatured")
	resp, err := just.GET(url, nil)
	if err != nil {
		return nil, err
	}

	var obj FeaturedMedia
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
