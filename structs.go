package gopenai

import (
	"net/http"
)

var availableModels *ListModels

type Client struct {
	basicAuth basicAuth
	client    *http.Client
	SubClient interface{}
}

type basicAuth struct {
	apiKey       string
	organisation string
}

type Dalle struct {
	client    *http.Client
	basicAuth basicAuth
}
