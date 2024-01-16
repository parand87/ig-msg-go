package instagram

import (
	"encoding/json"
	"github.com/parand87/ig-msg-go/instagram/fields"
	"net/url"
	"strings"
	"time"
)

type Message struct {
	Id          string    `json:"id,omitempty"`
	Message     string    `json:"message,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
}

type MessageListResponse struct {
	Messages ListResponse[Message] `json:"messages,omitempty"`
}

var messageFields = []string{
	fields.Id,
	fields.Message,
	fields.CreatedTime,
	fields.Reactions,
	fields.Shares,
}

func (i *Instagram) GetMessages(pageAccessToken string, conversationId string) ([]Message, error) {
	params := url.Values{}
	params.Set(fields.Fields, "messages")
	if pageAccessToken != "" {
		params.Set(fields.AccessToken, pageAccessToken)
	}
	endpoint := i.Config.Domain + "/" + conversationId + "?" + params.Encode()
	messageResponse, err := sendRequest[MessageListResponse](endpoint)

	var messages []Message

	for _, message := range messageResponse.Messages.Data {
		fullMessage, err := i.GetMessage(pageAccessToken, message.Id)
		if err != nil {
			return nil, err
		}
		messages = append(messages, fullMessage)
	}
	return messages, err
}

func (i *Instagram) GetMessage(pageAccessToken string, messageId string) (Message, error) {
	params := url.Values{}
	params.Set(fields.Fields, strings.Join(messageFields, ","))
	if pageAccessToken != "" {
		params.Set(fields.AccessToken, pageAccessToken)
	}
	endpoint := i.Config.Domain + "/" + messageId + "?" + params.Encode()
	message, err := sendRequest[Message](endpoint)
	return message, err
}

func (c *Message) UnmarshalJSON(b []byte) error {
	type Alias Message
	aux := &struct {
		CreatedTime string `json:"created_time,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	// Parse the UpdatedTime field with the specified layout
	parsedTime, err := time.Parse("2006-01-02T15:04:05-0700", aux.CreatedTime)
	if err != nil {
		return err
	}

	c.CreatedTime = parsedTime

	return nil
}
