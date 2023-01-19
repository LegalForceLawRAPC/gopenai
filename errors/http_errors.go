package errors

import (
	"fmt"
	"net/http"
)

type httpError error

type HttpError Error

var (
	ErrInvalidApiKey = HttpError{
		Error: httpError(fmt.Errorf("invalid api key")),
		Code:  "HTTP01",
	}
	ErrDefault = HttpError{
		Error: httpError(fmt.Errorf("an error occurred")),
		Code:  "HTTP02",
	}
)

func HandleHttpError(err error, statusCode int) *HttpError {
	switch statusCode {
	case http.StatusUnauthorized:
		return &ErrInvalidApiKey
	case http.StatusNotFound:
		return &ErrInvalidApiKey
	default:
		return &ErrDefault
	}
}
