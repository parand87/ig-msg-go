package instagram

import (
	"net/url"
	"strings"
)

func (i *Instagram) GetFacebookLoginUrl(permissions []string) string {
	var params = url.Values{}
	params.Set("client_id", i.Config.AppId)
	params.Set("redirect_uri", i.Config.RedirectUri)
	params.Set("response_type", i.Config.ResponseType)
	params.Set("scope", strings.Join(permissions, ","))
	return i.Config.FBLoginDomain + i.Config.Prefix + "/dialog/oauth" + "?" + params.Encode()
}
