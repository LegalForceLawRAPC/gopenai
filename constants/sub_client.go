package constants

import "net/http"

type SubClient struct {
	Client    *http.Client
	BasicAuth basicAuth
}

type basicAuth struct {
	ApiKey       string
	Organisation string
}
