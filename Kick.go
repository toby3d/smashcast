package hitGox

import (
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// KickTeamMember kick a team member.
//
// Team creators that want to leave must disband the team.
func (account *Account) KickTeamMember(team *Team, user string) (*just.Status, error) {
	var args f.Args
	args.Add("authToken", account.AuthToken)
	args.Add("group_id", team.Info.GroupID)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("team/", team.Info.GroupName, "/", user))
	resp, err := just.DELETE(url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
