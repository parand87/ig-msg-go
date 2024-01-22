package instagram

import (
	"github.com/parand87/ig-msg-go/instagram/constants"
	"net/url"
	"strings"
)

type Account struct {
	Id                       string          `json:"id,omitempty"`
	Name                     string          `json:"name,omitempty"`
	Username                 string          `json:"username,omitempty"`
	AccessToken              string          `json:"access_token,omitempty"`
	InstagramBusinessAccount BusinessAccount `json:"instagram_business_account"`
}

type BusinessAccount struct {
	Id string `json:"id"`
}

const AccountsEndpoint = "/me/accounts"

var accountFields = []string{
	constants.Fields.AccessToken,
	constants.Fields.Bio,
	constants.Fields.Id,
	constants.Fields.Name,
	constants.Fields.Username,
	constants.Fields.InstagramBusinessAccount,
}

func (i *Instagram) GetAccounts() ([]Account, error) {
	var params = url.Values{}
	params.Set(constants.Fields.Fields, strings.Join(accountFields, ","))
	if i.Config.AccessToken != "" {
		params.Set(constants.Fields.AccessToken, i.Config.AccessToken)
	}

	endpoint := i.Config.Domain + AccountsEndpoint + "?" + params.Encode()
	data, err := sendRequest[ListResponse[Account]](endpoint)
	return data.Data, err
}

func (i *Instagram) GetPageAccessToken(accountId string) (*Account, error) {
	var params = url.Values{}
	params.Set(constants.Fields.Fields, strings.Join(accountFields, ","))
	if i.Config.AccessToken != "" {
		params.Set(constants.Fields.AccessToken, i.Config.AccessToken)
	}

	endpoint := i.Config.Domain + "/" + accountId + "?" + params.Encode()
	account, err := sendRequest[Account](endpoint)
	if err != nil {
		return nil, err
	}

	i.SetPageAccessToken(account.AccessToken)
	return &account, err
}
