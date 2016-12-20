package hitGox

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

// CheckToken checks if the Token is valid.
func (app *Application) CheckToken(authToken string) (*Status, error) {
	if err := checkCheckToken(app, authToken); err != nil {
		return nil, err
	}

	var args fasthttp.Args
	args.Add("token", authToken)

	url := fmt.Sprint(API, "/auth/valid/", app.Name)
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	status, err := fuckYouNeedDecodeStatusFirst(resp)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func checkCheckToken(app *Application, authToken string) error {
	switch {
	case app.Name == "":
		return errors.New("no name of application, create new application first")
	case authToken == "":
		return errors.New("authtoken can not be empty")
	}
	return nil
}
