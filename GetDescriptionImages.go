package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
)

// DescriptionImage is a single element of list abot uploaded description image.
type DescriptionImage []struct {
	ImageID        string `json:"image_id"`
	ImagePath      string `json:"image_path"`
	ImageDateAdded string `json:"image_date_added"`
}

// GetDescriptionImages gets description images.
func (account *Account) GetDescriptionImages() (DescriptionImage, error) {
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("upload/description/", account.UserName, "/", account.AuthToken))
	resp, err := just.GET(url, nil)
	if err != nil {
		return nil, err
	}

	var obj DescriptionImage
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj, nil
}
