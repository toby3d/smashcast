package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
)

// CreateTeam creates a team object.
//
// displayName must match dame except casing.
func (account *Account) CreateTeam(name string, displayName string, text string) (*Status, error) {
	switch {
	case account.AuthToken == "":
		return nil, errors.New("authtoken in account can not be empty")
	case account.UserName == "":
		return nil, errors.New("username in account can not be empty")
	case name == "" || len(name) < 4:
		return nil, errors.New("name too short")
	case text == "" || len(text) < 4:
		return nil, errors.New("text required")
	case displayName == "":
		return nil, errors.New("invalid display name")
	}

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

	var args fasthttp.Args
	args.Add("authToken", account.AuthToken)

	url := APIEndpoint + "/team"
	resp, err := post(dst, url, &args)
	if err != nil {
		return nil, err
	}

	var obj Status
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
