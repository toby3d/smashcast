package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// SubscriptionInfo is about user subscription status.
type SubscriptionInfo struct {
	SubID            string   `json:"sub_id"`
	SubDateAdded     string   `json:"sub_date_added"`
	SubDateValid     string   `json:"sub_date_valid"`
	SubPlanID        string   `json:"sub_plan_id"`
	SubPaymentMethod string   `json:"sub_payment_method"`
	PlanCharge       string   `json:"plan_charge"`
	PlanCurrency     string   `json:"plan_currency"`
	PlanRecurring    string   `json:"plan_recurring"`
	UserName         string   `json:"user_name"`
	UserID           string   `json:"user_id"`
	UserLogo         string   `json:"user_logo"`
	UserLogoSmall    string   `json:"user_logo_small"`
	Cancel           string   `json:"cancel"`
	Benefits         []string `json:"benefits"`
	Resub            bool     `json:"resub"`
}

// CheckSubscriptionInfo retruns subscription information between :channel and :user
func (account *Account) CheckSubscriptionInfo(user string) (*SubscriptionInfo, error) {
	var args f.Args
	args.Add("authToken", account.AuthToken)

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("subscription/", account.UserName, "/", user))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj SubscriptionInfo
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
