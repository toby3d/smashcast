package hitGox

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// RemoveDescriptionImage removes uploaded description image.
func (account *Account) RemoveDescriptionImage(imageID interface{}) (*Status, error) {
	if err := checkRemoveDescriptionImage(account, imageID); err != nil {
		return nil, err
	}

	var args fasthttp.Args
	switch id := imageID.(type) {
	case int:
		args.Add("image_id", strconv.Itoa(id))
	case string:
		args.Add("image_id", id)
	default:
		return nil, errors.New("id type must be only as string or int")
	}

	url := fmt.Sprint(API, "/upload/description/", account.UserName, "/", account.AuthToken)
	resp, err := delete(url, &args)
	if err != nil {
		return nil, err
	}

	status, err := fuckYouNeedDecodeStatusFirst(resp)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func checkRemoveDescriptionImage(account *Account, imageID interface{}) error {
	switch {
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case account.UserName == "":
		return errors.New("username in account can not be empty")
	case imageID == nil:
		return errors.New("imageid can not be empty")
	}
	return nil
}
