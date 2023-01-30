package dalle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/LegalForceLawRAPC/gopenai/constants"
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
	r := &EditImageResponse{}
	var (
		buf = new(bytes.Buffer)
		w   = multipart.NewWriter(buf)
	)
	// If image is not nil, add it to the request
	if e.Image != nil {
		part, err := w.CreateFormFile("image", "/img/image")
		if err != nil {
			// multipart form error
			return nil, err
		}
		_, err = part.Write(e.Image)
		if err != nil {
			// multipart form error
			return nil, err
		}
	}
	// If mask is not nil, add it to the request
	if e.Mask != nil {
		part, err := w.CreateFormFile("mask", "/img/mask")
		if err != nil {
			// multipart form error
			return nil, err
		}
		_, err = part.Write(e.Mask)
		if err != nil {
			// multipart form error
			return nil, err
		}
	}
	// Add the json data to the request
	prompt := e.req.Body.(EditImagesPrompt).Prompt
	n := e.req.Body.(EditImagesPrompt).N
	size := e.req.Body.(EditImagesPrompt).Size
	err := w.WriteField("prompt", prompt)
	if err != nil {
		return nil, err
	}
	err = w.WriteField("n", fmt.Sprintf("%d", n))
	if err != nil {
		return nil, err
	}
	err = w.WriteField("size", size)
	if err != nil {
		return nil, err
	}
	if err != nil {
		// multipart form error
		return nil, err
	}
	request, err := http.NewRequest(e.req.Method, fmt.Sprintf("%s/%s", constants.BaseURL, e.req.Endpoint), buf)
	if err != nil {
		// http request error
		return nil, err
	}
	// Adding the bearer token header
	request.Header.Set("Authorization", constants.GetToken())
	// Now we add the multipart file data
	// The image and mask are added as multipart form data
	request.Header.Set("Content-Type", w.FormDataContentType())
	// Send the request
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		// http request error
		return nil, err
	}
	// Read the response
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		// io error
		return nil, err
	}
	// Unmarshal the response into the EditImageResponse struct
	err = json.Unmarshal(respBodyBytes, r)
	if err != nil {
		// json unmarshalling error
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		// The request was not successful
		return r, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	return r, nil
}
