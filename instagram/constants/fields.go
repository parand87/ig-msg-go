package constants

type fieldsConstants struct {
	AccessToken  string
	Fields       string
	Platform     string
	Bio          string
	Id           string
	Name         string
	Username     string
	Participants string
	UpdatedTime  string
	UnreadCount  string
	Reactions    string
	Shares       string
	Message      string
	CreatedTime  string
}

var Fields = fieldsConstants{
	AccessToken:  "access_token",
	Fields:       "fields",
	Platform:     "platform",
	Bio:          "bio",
	Id:           "id",
	Name:         "name",
	Username:     "username",
	Participants: "participants",
	UpdatedTime:  "updated_time",
	UnreadCount:  "unread_count",
	Reactions:    "reactions",
	Shares:       "shares",
	Message:      "message",
	CreatedTime:  "created_time",
}
