package hitGox

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// AcceptTeamInvite accept an invite from teamName.
func (account *Account) AcceptTeamInvite(teamName string, groupID interface{}) (*Status, error) {
	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)
	switch id := groupID.(type) {
	case int:
		args.Add("group_id", strconv.Itoa(id))
	case string:
		args.Add("group_id", id)
	default:
		return nil, errors.New("groupid can be only as string or int")
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("/team/", teamName, "/", account.UserName))
	resp, err := update(url, &args)
	if err != nil {
		return nil, err
	}

	return fixStatus(resp), nil
}
