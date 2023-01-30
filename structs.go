package gopenai

import (
	"net/http"
)

var availableModels *ListModels

// Client is the main struct for the gopenai package
type Client struct {
	basicAuth basicAuth
	client    *http.Client
	SubClient interface{}
}

// basicAuth is a struct that holds the api key and organisation id
type basicAuth struct {
	apiKey       string
	organisation string
}
