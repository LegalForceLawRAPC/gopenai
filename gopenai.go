package gopenai

import (
	"net/http"
	"time"

	"github.com/LegalForceLawRAPC/gopenai/constants"
	"github.com/LegalForceLawRAPC/gopenai/dalle"
)

var c *http.Client

// NewClient creates a new openAI client
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
	constants.SetApiKey(apiKey)
	constants.SetOrgId(organisation)
	c.basicAuth.apiKey = apiKey
	c.basicAuth.organisation = organisation
	l := &ListModels{}
	err := c.Do(constants.GetOpenAIEndpoint("listModels"), &l)
	if err != nil {
		return err.Error
	}
	return nil
}

// ListModels returns a list of available models
func (c *Client) ListModels() *ListModels {
	if availableModels != nil {
		return availableModels
	} else {
		l := &ListModels{}
		err := c.Do(constants.GetOpenAIEndpoint("listModels"), &l)
		if err != nil {
			return nil
		}
		availableModels = l
		return l
	}
}

// Dalle returns a new dalle client
func (c *Client) Dalle() *dalle.Dalle {
	return &dalle.Dalle{
		Client: c.client,
		BasicAuth: constants.BasicAuth{
			ApiKey:       c.basicAuth.apiKey,
			Organisation: c.basicAuth.organisation,
		},
	}
}
