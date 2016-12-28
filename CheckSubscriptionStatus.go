package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CheckSubscriptionStatus rReturns subscription relationship between channel and auth.
func (account *Account) CheckSubscriptionStatus(channel string) (bool, error) {
	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("user/subscription/", channel, "/", account.AuthToken))
	resp, err := get(url, nil)
	if err != nil {
		return false, err
	}

	var obj = struct {
		IsSubscriber bool `json:"isSubscriber"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj.IsSubscriber, nil
}
