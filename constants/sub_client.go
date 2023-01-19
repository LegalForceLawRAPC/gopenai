package constants

import "net/http"

type SubClient struct {
	Client    *http.Client
	BasicAuth BasicAuth
}

type BasicAuth struct {
	ApiKey       string
	Organisation string
}
