package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// DescriptionImage is a single element of list abot uploaded description image.
type DescriptionImage []struct {
	ImageID        string `json:"image_id"`
	ImagePath      string `json:"image_path"`
	ImageDateAdded string `json:"image_date_added"`
}

// GetDescriptionImages gets description images.
func (account *Account) GetDescriptionImages() (DescriptionImage, error) {
	if err := checkGetDescriptionImages(account); err != nil {
		return nil, err
	}

	url := fmt.Sprint(API, "/upload/description/", account.UserName, "/", account.AuthToken)
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj DescriptionImage
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func checkGetDescriptionImages(account *Account) error {
	switch {
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case account.UserName == "":
		return errors.New("username in account can not be empty")
	}
	return nil
}
