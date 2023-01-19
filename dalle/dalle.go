package dalle

import (
	"github.com/LegalForceLawRAPC/gopenai/constants"
	"io"
	"log"
	"os"
)

func (d *Dalle) GenerateImages(prompt string, n int, size string, userId string) (*GenerateImagesResponse, error) {
	r := &GenerateImagesResponse{}
	req := constants.GetDalleEndpoint("generateImages")
	req.Body = GenerateImagesPrompt{
		Prompt:   prompt,
		N:        n,
		Size:     size,
		UserName: userId,
	}
	err := d.Do(req, &r)
	if err != nil {
		return nil, err.Error
	}
	return r, nil
}

func (d *Dalle) InitEditImages(prompt string, n int, size string, userId string) *EditImagesRequest {
	req := constants.GetDalleEndpoint("editImages")
	body := EditImagesPrompt{
		Prompt:   prompt,
		N:        n,
		Size:     size,
		UserName: userId,
	}
	req.Body = body
	return &EditImagesRequest{
		req: req,
	}
}

func (e *EditImagesRequest) AddImageFromFileSystem(image *os.File) *EditImagesRequest {
	im, err := io.ReadAll(image)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	e.Image = im
	return e
}

func (e *EditImagesRequest) AddMaskFromFileSystem(mask *os.File) *EditImagesRequest {
	im, err := io.ReadAll(mask)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	e.Mask = im
	return e
}

func (e *EditImagesRequest) AddImage(image []byte) *EditImagesRequest {
	e.Image = image
	return e
}

func (e *EditImagesRequest) AddMask(mask []byte) *EditImagesRequest {
	e.Mask = mask
	return e
}

func (e *EditImagesRequest) Do() (*EditImageResponse, error) {
	//res := &EditImageResponse{}
	//body := &bytes2.Buffer{}
	//writer := multipart.NewWriter(body)
	//part, err := writer.CreateFormFile("image", "image.png")
	//if err != nil {
	//	return nil, err
	//}
	//req, err := http.NewRequest(e.req.Method, fmt.Sprintf("%s/%s", constants.BaseURL, e.req.Endpoint), bytes2.NewBuffer(m))
	//if err != nil {
	//	return nil, err
	//}
	//req.Header.Set("Authorization", constants.GetToken())
	return nil, nil

}
