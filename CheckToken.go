package hitGox

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// CheckToken checks if the Token is valid.
func (app *OAuthApplication) CheckToken(authToken string) (*Status, error) {
	var args fasthttp.Args
	args.Add("token", authToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("auth/valid/", app.Name))
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	return fixStatus(resp), nil
}
