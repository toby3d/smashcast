package hitGox

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// SetDefaultTeam sets default team by id.
func (account *Account) SetDefaultTeam(id int) (*Status, error) {
	var changes = struct {
		GroupID string `json:"group_id"`
	}{strconv.Itoa(id)}

	dst, err := json.Marshal(changes)
	if err != nil {
		return nil, err
	}

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("user/", account.UserName, "/team/default"))
	resp, err := post(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return fixStatus(resp), nil
}
