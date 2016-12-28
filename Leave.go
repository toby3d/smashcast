package hitGox

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// LeaveFromTeam leave the team.
//
// Team creators that want to leave must disband the team.
func (account *Account) LeaveFromTeam(team *Team) (*Status, error) {
	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)
	args.Add("group_id", team.Info.GroupID)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("team/", team.Info.GroupName, "/", account.UserName))
	resp, err := delete(url, &args)
	if err != nil {
		return nil, err
	}

	return fixStatus(resp), nil
}
