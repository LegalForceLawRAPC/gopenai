package dalle

import "github.com/LegalForceLawRAPC/gopenai/constants"

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
