package gopenai

import (
	"net/http"
	"time"
)

var c *http.Client

func NewClient() *Client {
	c = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		},
	}
	return &Client{
		basicAuth: basicAuth{
			apiKey:       "",
			organisation: "",
		},
		client: c,
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
	return nil
}

func (c *Client) ListModels() *ListModels {
	l := &ListModels{}
	err := c.Do(openAiEndpoints["listModels"], &l)
	if err != nil {
		return nil
	}
	return l
}
