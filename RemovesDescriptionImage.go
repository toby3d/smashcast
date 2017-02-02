package hitGox

import (
	"fmt"
	"strconv"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// RemoveDescriptionImage removes uploaded description image.
func (account *Account) RemoveDescriptionImage(imageID interface{}) (*just.Status, error) {
	var args f.Args
	switch id := imageID.(type) {
	case int:
		args.Add("image_id", strconv.Itoa(id))
	case string:
		args.Add("image_id", id)
	default:
		return nil, fmt.Errorf("id type must be only as string or int")
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("upload/description/", account.UserName, "/", account.AuthToken))
	resp, err := just.DELETE(url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
