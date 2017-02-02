package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
)

type CommercialBreak struct {
	Method string `json:"method"`
	Params struct {
		Channel   string `json:"channel"`
		Error     string `json:"error"`
		Count     string `json:"count"`
		Delay     string `json:"delay"`
		Timestamp int    `json:"timestamp"`
		Token     string `json:"token"`
		URL       string `json:"url"`
	} `json:"params"`
}

func (account *Account) RunCommercial(channel string, adCount int) (*CommercialBreak, error) {
	var body = struct {
		UserName  string `json:"user_name"`
		AuthToken string `json:"authToken"`
	}{channel, account.AuthToken}

	dst, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprintf("ws/combreak/%s/%d", channel, adCount))
	resp, err := just.POST(dst, url, nil)
	if err != nil {
		return nil, err
	}

	var obj CommercialBreak
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
