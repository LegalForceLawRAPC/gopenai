package dalle

import (
	"github.com/LegalForceLawRAPC/gopenai/constants"
)

func (d *Dalle) GenerateImages(prompt string, n int, size string) (*GenerateImagesResponse, error) {
	r := &GenerateImagesResponse{}
	req := constants.GetDalleEndpoint("generateImages")
	req.Body = GenerateImagesPrompt{
		Prompt: prompt,
		N:      n,
		Size:   size,
	}
	err := d.Do(req, &r)
	if err != nil {
		return nil, err.Error
	}
	return r, nil
}
