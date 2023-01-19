package dalle

import "github.com/LegalForceLawRAPC/gopenai/constants"

type Dalle constants.SubClient

type GenerateImagesPrompt struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}
