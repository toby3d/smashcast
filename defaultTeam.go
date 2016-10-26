package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// DefaultTeam sets default team by groupID.
func DefaultTeam(user UserName, groupID int, token Token) (Status, error) {
	args := fasthttp.AcquireArgs()
	args.Add("authToken", token.Token)
	args.Add("group_id", strconv.Itoa(groupID))
	requestURL := fmt.Sprintf("%s/user/%s/team/default", API, user.UserName)
	statusCode, body, err := fasthttp.Post(nil, requestURL, args)
	if statusCode != 200 || err != nil {
		return Status{}, err
	}
	var obj Status
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Status{}, err
	}
	return obj, nil
}
