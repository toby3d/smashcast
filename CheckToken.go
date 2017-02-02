package hitGox

import (
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// CheckToken checks if the Token is valid.
func (app *OAuthApplication) CheckToken(authToken string) (*just.Status, error) {
	var args f.Args
	args.Add("token", authToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("auth/valid/", app.Name))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
