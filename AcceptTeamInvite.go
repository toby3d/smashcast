package hitGox

import (
	"fmt"
	"strconv"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// AcceptTeamInvite accept an invite from teamName.
func (account *Account) AcceptTeamInvite(teamName string, groupID interface{}) (*just.Status, error) {
	var args f.Args
	args.Add("authToken", account.AuthToken)
	switch id := groupID.(type) {
	case int:
		args.Add("group_id", strconv.Itoa(id))
	case string:
		args.Add("group_id", id)
	default:
		return nil, fmt.Errorf("groupid can be only as string or int")
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("team/", teamName, "/", account.UserName))
	resp, err := just.UPDATE(url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
