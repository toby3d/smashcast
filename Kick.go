package hitGox

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// KickTeamMember kick a team member.
//
// Team creators that want to leave must disband the team.
func (account *Account) KickTeamMember(team *Team, user string) (*Status, error) {
	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)
	args.Add("group_id", team.Info.GroupID)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("team/", team.Info.GroupName, "/", user))
	resp, err := delete(url, &args)
	if err != nil {
		return nil, err
	}

	return fixStatus(resp), nil
}
