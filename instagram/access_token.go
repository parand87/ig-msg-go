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

type DebugToken struct {
	AppId          string  `json:"app_id,omitempty"`
	ExpiresAt      int64   `json:"expires_at,omitempty"`
	IsValid        bool    `json:"is_valid,omitempty"`
	UserId         string  `json:"user_id,omitempty"`
	ProfileId      string  `json:"profile_id,omitempty"`
	GranularScopes []Scope `json:"granular_scopes,omitempty"`
	Type           string  `json:"type"`
}
type AccessTokenType string

type Scope struct {
	Scope     string   `json:"scope,omitempty"`
	TargetIds []string `json:"target_ids,omitempty"`
}

const (
	PageToken AccessTokenType = "PAGE"
	UserToken AccessTokenType = "USER"
)

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

func (i *Instagram) ValidateToken(userToken string, inputToken string) (DebugToken, error) {
	var params = url.Values{}
	params.Set("access_token", userToken)
	params.Set("input_token", inputToken)

	var debugTokenUri = i.Config.Domain + "/debug_token" + "?" + params.Encode()

	data, err := sendRequest[Response[DebugToken]](debugTokenUri)
	return data.Data, err
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

	return data, err
}
func (a *AccessToken) IsLongLived() bool {
	return true
}
