package hitGox

import (
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// CreateTeam creates a team object.
//
// displayName must match dame except casing.
func (account *Account) CreateTeam(name string, displayName string, text string) (*just.Status, error) {
	var changes = struct {
		AuthToken        string `json:"authToken"`
		GroupUserName    string `json:"group_user_name"`
		GroupName        string `json:"group_name"`
		GroupText        string `json:"group_text"`
		GroupDisplayName string `json:"group_display_name"`
	}{account.AuthToken, account.UserName, name, text, displayName}

	dst, err := json.Marshal(changes)
	if err != nil {
		return nil, err
	}

	var args f.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, "team")
	resp, err := just.POST(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
