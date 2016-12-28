package hitGox

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// RemoveDescriptionImage removes uploaded description image.
func (account *Account) RemoveDescriptionImage(imageID interface{}) (*Status, error) {
	var args fasthttp.Args
	switch id := imageID.(type) {
	case int:
		args.Add("image_id", strconv.Itoa(id))
	case string:
		args.Add("image_id", id)
	default:
		return nil, errors.New("id type must be only as string or int")
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("upload/description/", account.UserName, "/", account.AuthToken))
	resp, err := delete(url, &args)
	if err != nil {
		return nil, err
	}

	return fixStatus(resp), nil
}
