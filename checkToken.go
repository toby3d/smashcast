package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// CheckToken checks if the authToken is valid.
func CheckToken(app Application, authToken AuthToken) (Status, error) {
	var args fasthttp.Args
	args.Add("authToken", authToken.Token)
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
