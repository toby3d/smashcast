package hitGox

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// AcceptTeamInvite accept an invite from teamName.
func (account *Account) AcceptTeamInvite(teamName string, groupID interface{}) (*Status, error) {
	if err := checkAcceptTeamInvite(account, teamName, groupID); err != nil {
		return nil, err
	}

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

	url := fmt.Sprint(API, "/team/", teamName, "/", account.UserName)
	resp, err := update(url, &args)
	if err != nil {
		return nil, err
	}

	status, err := fuckYouNeedDecodeStatusFirst(resp)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func checkAcceptTeamInvite(account *Account, teamName string, groupID interface{}) error {
	switch {
	case account.UserName == "":
		return errors.New("username in account can not be empty")
	case account.AuthToken == "":
		return errors.New("authtoken in account can not be empty")
	case teamName == "":
		return errors.New("teamname can not be empty")
	case groupID == nil:
		return errors.New("groupid can not be empty")
	}
	return nil

}
