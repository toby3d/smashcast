package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// DescriptionImage is a single element of list abot uploaded description image.
type DescriptionImage []struct {
	ImageID        string `json:"image_id"`
	ImagePath      string `json:"image_path"`
	ImageDateAdded string `json:"image_date_added"`
}

// GetDescriptionImages gets description images.
func (account *Account) GetDescriptionImages() (*[]DescriptionImage, error) {
	url := fmt.Sprint(API, "/upload/description/", account.UserName, "/", account.AuthToken)
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj []DescriptionImage
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
