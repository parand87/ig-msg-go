package instagram

import (
	"encoding/json"
	"github.com/parand87/ig-msg-go/instagram/fields"
	"net/url"
	"strings"
	"time"
)

const ConversationsEndpoint = "/me/conversations"

type Conversation struct {
	Id           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Participants []Participant `json:"participants,omitempty"`
	UpdatedTime  time.Time     `json:"updated_time,omitempty"`
	UnreadCount  int           `json:"unread_count,omitempty"`
}

type Participant struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

var conversationFields = []string{
	fields.Id,
	fields.Name,
	fields.Participants,
	fields.UpdatedTime,
	fields.UnreadCount,
}

func (i *Instagram) GetConversations(pageAccessToken string) ([]Conversation, error) {
	params := url.Values{}
	params.Set(fields.Fields, strings.Join(conversationFields, ","))
	params.Set(fields.Platform, "instagram")
	if pageAccessToken != "" {
		params.Set(fields.AccessToken, pageAccessToken)
	}
	endpoint := i.Config.Domain + ConversationsEndpoint + "?" + params.Encode()
	data, err := sendRequest[ListResponse[Conversation]](endpoint)
	return data.Data, err
}

func (c *Conversation) UnmarshalJSON(b []byte) error {
	type Alias Conversation
	aux := &struct {
		UpdatedTime  string `json:"updated_time,omitempty"`
		Participants struct {
			Data []Participant `json:"data,omitempty"`
		} `json:"participants,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	// Parse the UpdatedTime field with the specified layout
	parsedTime, err := time.Parse("2006-01-02T15:04:05-0700", aux.UpdatedTime)
	if err != nil {
		return err
	}

	c.UpdatedTime = parsedTime

	c.Participants = aux.Participants.Data
	return nil
}
