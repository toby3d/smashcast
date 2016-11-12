package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// StreamKey is a key for authenticate streaming programm.
type DescriptionImage struct {
	ImageID        string `json:"image_id,omitempty"`
	ImagePath      string `json:"image_path,omitempty"`
	ImageDateAdded string `json:"image_date_added,omitempty"`
}

// GetDescriptionImages gets description images.
func GetDescriptionImages(channel string, authToken AuthToken) ([]DescriptionImage, error) {
	var args fasthttp.Args
	args.Add("authToken", authToken.AuthToken)
	requestURL := fmt.Sprintf("%s/upload/description/%s/%s", API, channel, authToken.AuthToken)
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return nil, err
	}
	var obj []DescriptionImage
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return nil, err
	}
	return obj, nil
}
