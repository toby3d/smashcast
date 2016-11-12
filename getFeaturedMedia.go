package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// User contains information about user.
type FeaturedMedia struct {
	MediaID          string `json:"media_id,omitempty"`
	MediaDisplayName string `json:"media_display_name,omitempty"`
	MediaName        string `json:"media_name,omitempty"`
	Backdrop         string `json:"backdrop,omitempty"`
	BackdropHTML     string `json:"backdrop_html,omitempty"`
}

// GetFeaturedMedia return information about user.
//
// When a user isn’t found, this API returns a regular response but with all values containing null.
func GetFeaturedMedia() (FeaturedMedia, error) {
	_, body, err := fasthttp.Get(nil, API+"/mediafeatured")
	if err != nil {
		return FeaturedMedia{}, err
	}
	var obj FeaturedMedia
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return FeaturedMedia{}, err
	}
	return obj, nil
}
