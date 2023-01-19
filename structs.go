package gopenai

import "net/http"

type Client struct {
	basicAuth basicAuth
	client    *http.Client
}

type basicAuth struct {
	apiKey       string
	organisation string
}

type RequestData struct {
	endpoint string
	method   string
	body     interface{}
}
