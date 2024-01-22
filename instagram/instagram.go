package instagram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Cursor struct {
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}

type Paging struct {
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}

type ListResponse[T any] struct {
	Data   []T     `json:"data,omitempty"`
	Paging *Paging `json:"paging,omitempty"`
}
type Response[T any] struct {
	Data T `json:"data,omitempty"`
}

type Instagram struct {
	Config *Config
	Client *http.Client
}

func NewClient(config *Config) *Instagram {
	config.ResponseType = "code"
	config.GrantType = "authorization_code"
	config.FBLoginDomain = "https://www.facebook.com"
	config.Domain = "https://graph.facebook.com"
	config.Prefix = "/v18.0"

	var instagram = &Instagram{Config: config}
	instagram.Client = http.DefaultClient
	return instagram
}

func sendRequest[T any](endpoint string) (T, error) {
	var data T

	resp, err := http.Get(endpoint)
	if err != nil {
		return data, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err.Error())
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return data, err
	}
	return data, nil
}
