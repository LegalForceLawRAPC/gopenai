package gopenai

import (
	"log"
	"net/http"
)

func NewClient() *Client {
	return &Client{
		basicAuth: basicAuth{
			apiKey:       "",
			organisation: "",
		},
		client: &http.Client{},
	}
}

// Connect to the OpenAI API
func (c *Client) Connect(apiKey string, organisation string) error {
	c.basicAuth.apiKey = apiKey
	c.basicAuth.organisation = organisation
	l := &ListModels{}
	err := c.Do(openAiEndpoints["listModels"], &l)
	if err != nil {
		return err.Error
	}
	log.Println(l)
	return nil
}
