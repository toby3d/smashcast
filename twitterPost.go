package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// TwitterPost send Tweet To Twitter.
func TwitterPost(userName UserName, message string, token Token) (Status, error) {
	args := fasthttp.AcquireArgs()
	args.Add("user_name", userName.UserName)
	args.Add("authToken", token.Token)
	args.Add("message", message)
	statusCode, body, err := fasthttp.Post(nil, API+"/twitter/post", args)
	if statusCode != 200 || err != nil {
		return Status{}, err
	}
	var obj Status
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Status{}, err
	}
	return obj, nil
}
