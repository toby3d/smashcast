package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
)

// GetToken return authentication token rather than account information.
func (app *OAuthApplication) GetToken(login string, pass string) (string, error) {
	var changes = struct {
		Login string `json:"login"`
		Pass  string `json:"pass"`
		App   string `json:"app"`
	}{login, pass, app.Name}

	dst, err := json.Marshal(changes)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf(APIEndpoint, "auth/token")
	resp, err := just.POST(dst, url, nil)
	if err != nil {
		return "", err
	}

	var obj = struct {
		AuthToken string `json:"authToken"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj.AuthToken, nil
}
