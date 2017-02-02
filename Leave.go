package hitGox

import (
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// LeaveFromTeam leave the team.
//
// Team creators that want to leave must disband the team.
func (account *Account) LeaveFromTeam(team *Team) (*just.Status, error) {
	var args f.Args
	args.Add("authToken", account.AuthToken)
	args.Add("group_id", team.Info.GroupID)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("team/", team.Info.GroupName, "/", account.UserName))
	resp, err := just.DELETE(url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
