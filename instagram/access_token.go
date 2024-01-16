package instagram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const TokenEndpoint = "/oauth/access_token"

type AccessToken struct {
	TokenType   string        `json:"token_type,omitempty"`
	AccessToken string        `json:"access_token,omitempty"`
	Error       FacebookError `json:"error"`
}

func (i *Instagram) RequestAccessToken(code string) (*AccessToken, error) {
	accessToken, err := i.GetShortLivedAccessToken(code)
	if err != nil {
		return nil, err
	}
	return i.GetLongLivedAccessToken(accessToken)
}

func (i *Instagram) GetShortLivedAccessToken(code string) (*AccessToken, error) {
	var params = url.Values{}
	params.Set("client_id", i.Config.AppId)
	params.Set("client_secret", i.Config.AppSecret)
	params.Set("redirect_uri", i.Config.RedirectUri)
	params.Set("grant_type", i.Config.GrantType)
	params.Set("code", code)

	var accessTokenUrl = i.Config.Domain + TokenEndpoint + "?" + params.Encode()
	resp, err := http.Get(accessTokenUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data = new(AccessToken)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 400 {
		return nil, &data.Error
	}

	return data, nil
}

func (i *Instagram) GetLongLivedAccessToken(accessToken *AccessToken) (*AccessToken, error) {
	var params = url.Values{}
	params.Set("client_id", i.Config.AppId)
	params.Set("client_secret", i.Config.AppSecret)
	params.Set("redirect_uri", i.Config.RedirectUri)
	params.Set("grant_type", "fb_exchange_token")
	params.Set("fb_exchange_token", accessToken.AccessToken)

	var accessTokenUrl = i.Config.Domain + TokenEndpoint + "?" + params.Encode()
	resp, err := http.Get(accessTokenUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data = new(AccessToken)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 400 {
		return nil, &data.Error
	}

	err = i.SetUserAccessToken(data.AccessToken)
	return data, err
}
func (a *AccessToken) IsLongLived() bool {
	return true
}
