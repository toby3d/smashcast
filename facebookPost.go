package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// FacebookPost send Facebook Post to enabled facebook pages.
func (token Token) FacebookPost(userName UserName, message string) (Status, error) {
	args := fasthttp.AcquireArgs()
	args.Add("user_name", userName.UserName)
	args.Add("authToken", token.Token)
	args.Add("message", message)
	statusCode, body, err := fasthttp.Post(nil, API+"/facebook/post", args)
	if statusCode != 200 || err != nil {
		return Status{}, err
	}
	var obj Status
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Status{}, err
	}
	return obj, nil
}
