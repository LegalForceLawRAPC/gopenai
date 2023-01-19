package dalle

import (
	"github.com/LegalForceLawRAPC/gopenai/constants"
)

type Dalle constants.SubClient

type GenerateImagesPrompt struct {
	Prompt   string `json:"prompt"`
	N        int    `json:"n"`
	Size     string `json:"size"`
	UserName string `json:"user"`
}

type GenerateImagesResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

type EditImagesPrompt struct {
	Prompt   string `json:"prompt"`
	N        int    `json:"n"`
	Size     string `json:"size"`
	UserName string `json:"user"`
}

type EditImagesRequest struct {
	Image []byte `json:"image"`
	Mask  []byte `json:"mask"`
	req   constants.RequestData
}

type EditImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}

type AddFiles struct {
	Image    []byte `json:"image"`
	FileName string `json:"file_name"`
}
