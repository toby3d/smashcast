package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// CheckToken checks if the `token` is valid.
func (app Application) CheckToken(token Token) (Status, error) {
	var args fasthttp.Args
	args.Add("token", token.Token)
	requestURL := fmt.Sprintf("%s/auth/valid/%s?%s", API, app.Name, args.String())
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return Status{}, err
	}
	var obj Status
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Status{}, err
	}
	return obj, nil
}
