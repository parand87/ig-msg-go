package instagram

import (
	"github.com/parand87/ig-msg-go/instagram/fields"
	"net/url"
	"strings"
)

type Account struct {
	Id             string `json:"id,omitempty"`
	FollowersCount int    `json:"followers_count,omitempty"`
	Name           string `json:"name,omitempty"`
	Username       string `json:"username,omitempty"`
	AccessToken    string `json:"access_token,omitempty"`
}

const AccountsEndpoint = "/me/accounts"

var accountFields = []string{
	fields.AccessToken,
	fields.Bio,
	fields.Id,
	fields.Name,
	fields.Username,
}

func (i *Instagram) GetAccounts() ([]Account, error) {
	var params = url.Values{}
	params.Set(fields.Fields, strings.Join(accountFields, ","))
	if i.Config.AccessToken != "" {
		params.Set(fields.AccessToken, i.Config.AccessToken)
	}

	endpoint := i.Config.Domain + AccountsEndpoint + "?" + params.Encode()
	data, err := sendRequest[ListResponse[Account]](endpoint)
	return data.Data, err
}

func (i *Instagram) GetPageAccessToken(accountId string) (*Account, error) {
	var params = url.Values{}
	params.Set(fields.Fields, strings.Join(accountFields, ","))
	if i.Config.AccessToken != "" {
		params.Set(fields.AccessToken, i.Config.AccessToken)
	}

	endpoint := i.Config.Domain + "/" + accountId + "?" + params.Encode()
	account, err := sendRequest[Account](endpoint)
	if err != nil {
		return nil, err
	}

	i.SetPageAccessToken(account.AccessToken)
	return &account, err
}
