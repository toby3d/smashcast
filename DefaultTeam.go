package hitGox

import (
	"encoding/json"
	"fmt"
	"strconv"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// SetDefaultTeam sets default team by id.
func (account *Account) SetDefaultTeam(id int) (*just.Status, error) {
	var changes = struct {
		GroupID string `json:"group_id"`
	}{strconv.Itoa(id)}

	dst, err := json.Marshal(changes)
	if err != nil {
		return nil, err
	}

	var args f.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("user/", account.UserName, "/team/default"))
	resp, err := just.POST(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
