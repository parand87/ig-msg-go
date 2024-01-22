package instagram

type UserData struct {
	Id          string `json:"id"`
	InstagramId string `json:"instagram_id"`
	UserToken   string `json:"user_token"`
	PageToken   string `json:"page_token"`
}
