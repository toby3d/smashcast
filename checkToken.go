package hitGox

import (
	"errors"
	"github.com/valyala/fasthttp"
)

// CheckToken checks if the Token is valid.
func (app *Application) CheckToken(authToken string) (*Status, error) {
	switch {
	case app.Name == "":
		return nil, errors.New("no name of application, create new application first")
	case authToken == "":
		return nil, errors.New("authtoken can not be empty")
	}

	var args fasthttp.Args
	args.Add("token", authToken)

	url := APIEndpoint + "/auth/valid/" + app.Name
	body, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	status, err := stupidFuckingStatusResponseByLazyAPIDevelopers(&body)
	if err != nil {
		return nil, err
	}

	return status, nil
}
