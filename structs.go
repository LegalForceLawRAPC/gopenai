package gopenai

import (
	"github.com/LegalForceLawRAPC/gopenai/constants"
	"net/http"
)

type Client struct {
	basicAuth basicAuth
	client    *http.Client
	subClient constants.SubClient
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
