package instagram

import (
	"bytes"
	"encoding/json"
	"github.com/parand87/ig-msg-go/instagram/constants"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Message struct {
	Id          string    `json:"id,omitempty"`
	Message     string    `json:"message,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
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

type MessageListResponse struct {
	Messages ListResponse[Message] `json:"messages,omitempty"`
}

var messageFields = []string{
	constants.Fields.Id,
	constants.Fields.Message,
	constants.Fields.CreatedTime,
	constants.Fields.Reactions,
	constants.Fields.Shares,
}

func (i *Instagram) GetMessages(conversationId string, userData *UserData) ([]Message, error) {
	params := url.Values{}
	params.Set(constants.Fields.Fields, "messages")
	params.Set(constants.Fields.AccessToken, userData.PageToken)
	endpoint := i.Config.Domain + "/" + conversationId + "?" + params.Encode()
	messageResponse, err := sendRequest[MessageListResponse](endpoint)

	var messages []Message

	for _, message := range messageResponse.Messages.Data {
		fullMessage, err := i.GetMessage(message.Id, userData)
		if err != nil {
			return nil, err
		}
		messages = append(messages, fullMessage)
	}
	return messages, err
}

func (i *Instagram) GetMessage(messageId string, userData *UserData) (Message, error) {
	params := url.Values{}
	params.Set(constants.Fields.Fields, strings.Join(messageFields, ","))
	params.Set(constants.Fields.AccessToken, userData.PageToken)
	endpoint := i.Config.Domain + "/" + messageId + "?" + params.Encode()
	message, err := sendRequest[Message](endpoint)
	return message, err
}

type MessageRecipient struct {
	Id string `json:"id"`
}
type MessageText struct {
	Text string `json:"text"`
}

func (i *Instagram) SendTextMessage(recipientId string, text string, userData *UserData) error {
	params := url.Values{}
	params.Set(constants.Fields.AccessToken, userData.PageToken)
	endpoint := i.Config.Domain + "/me/messages" + "?" + params.Encode()

	messageData := MessageRecipient{
		Id: recipientId,
	}
	recipientJson, err := json.Marshal(messageData)
	if err != nil {
		return err
	}
	textData := MessageText{
		Text: text,
	}
	textJson, err := json.Marshal(textData)
	if err != nil {
		return err
	}

	formData := url.Values{}
	formData.Set("recipient", string(recipientJson))
	formData.Set("message", string(textJson))
	formEncoded := formData.Encode()

	buffer := bytes.NewBufferString(formEncoded)
	response, err := http.Post(endpoint, "application/x-www-form-urlencoded", buffer)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err.Error())
		}
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		var fbError = new(FacebookErrorResponse)
		err = json.Unmarshal(body, fbError)
		if err != nil {
			return err
		}
		return &fbError.Error
	}

	return nil
}
