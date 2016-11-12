package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// CreateTeam get an authentication authToken rather than account information.
func (authToken AuthToken) CreateTeam(groupUserName string, groupName string, groupText string, groupDisplayName string) (Status, error) {
	args := fasthttp.AcquireArgs()
	args.Add("authToken", authToken.AuthToken)
	args.Add("group_user_name", groupUserName)
	args.Add("group_name", groupName)
	args.Add("group_text", groupText)

	// group_display_name must match group_name except casing.
	args.Add("group_display_name", groupDisplayName)

	statusCode, body, err := fasthttp.Post(nil, API+"/team", args)
	if statusCode != 200 || err != nil {
		return Status{}, err
	}
	var obj Status
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Status{}, err
	}
	return obj, nil
}
