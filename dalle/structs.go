package dalle

import (
	"github.com/LegalForceLawRAPC/gopenai/constants"
)

// Dalle allows access to all Dalle type endpoints
type Dalle constants.SubClient

// GenerateImagesPrompt is the prompt for the generateImages endpoint
type GenerateImagesPrompt struct {
	Prompt   string `json:"prompt"`
	N        int    `json:"n"`
	Size     string `json:"size"`
	UserName string `json:"user"`
}

// GenerateImagesRequest is the request for the generateImages endpoint
type GenerateImagesResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

// EditImagesPrompt is the prompt for the editImages endpoint
type EditImagesPrompt struct {
	Prompt   string `json:"prompt"`
	N        int    `json:"n"`
	Size     string `json:"size"`
	UserName string `json:"user"`
}

// EditImagesRequest is the request for the editImages endpoint
type EditImagesRequest struct {
	Image []byte `json:"image"`
	Mask  []byte `json:"mask"`
	req   constants.RequestData
}

// EditImageResponse is the response sent by the editImages endpoint
type EditImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

// AddFiles is the request for the addFiles endpoint
type AddFiles struct {
	Image    []byte `json:"image"`
	FileName string `json:"file_name"`
}
