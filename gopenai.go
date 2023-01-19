package gopenai

import (
	"github.com/LegalForceLawRAPC/gopenai/constants"
	"github.com/LegalForceLawRAPC/gopenai/dalle"
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
	err := c.Do(constants.GetOpenAIEndpoint("listModels"), &l)
	if err != nil {
		return err.Error
	}
	return nil
}

func (c *Client) ListModels() *ListModels {
	l := &ListModels{}
	err := c.Do(constants.GetOpenAIEndpoint("listModels"), &l)
	if err != nil {
		return nil
	}
	return l
}

func (c *Client) Dalle() *dalle.Dalle {
	return &dalle.Dalle{
		Client: c.client,
		BasicAuth: constants.BasicAuth{
			ApiKey:       c.basicAuth.apiKey,
			Organisation: c.basicAuth.organisation,
		},
	}
}
