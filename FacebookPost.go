package hitGox

import (
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// SendFacebookPost send Facebook post to enabled facebook pages.
func (account *Account) SendFacebookPost(message string) (*just.Status, error) {
	var body = struct {
		UserName  string `json:"user_name"`
		AuthToken string `json:"authToken"`
		Message   string `json:"message"`
	}{account.UserName, account.AuthToken, message}

	dst, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	var args f.Args
	args.Add("authToken", account.AuthToken)
	args.Add("user_name", account.UserName)

	url := fmt.Sprintf(APIEndpoint, "facebook/post")
	resp, err := just.POST(dst, url, &args)
	if err != nil {
		return nil, err
	}

	return just.FixStatus(resp), nil
}
