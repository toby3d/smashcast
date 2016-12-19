package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

	url := fmt.Sprint(API, "/auth/valid/", app.Name)
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	var obj Status
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
