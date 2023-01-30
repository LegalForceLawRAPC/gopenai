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

// GenerateImages generates images from a prompt
//
// prompt: The prompt to generate images from
// n: The number of images to generate
// size: The size of the images to generate
// userId: The user id to generate the images for that specific user, openAI keeps a track of the images generated for each user
func (d *Dalle) GenerateImages(prompt string, n int, size string, userId string) (*GenerateImagesResponse, error) {
	// Create the response struct
	r := &GenerateImagesResponse{}
	// Get the request Data
	req := constants.GetDalleEndpoint("generateImages")
	// Initialise the request body
	req.Body = GenerateImagesPrompt{
		Prompt:   prompt,
		N:        n,
		Size:     size,
		UserName: userId,
	}
	// Send the request
	err := d.Do(req, &r)
	if err != nil {
		return nil, err.Error
	}
	// Return the response
	return r, nil
}

// InitEditImages initialises the EditImagesRequest struct
func (d *Dalle) InitEditImages(prompt string, n int, size string, userId string) *EditImagesRequest {
	// Get the request Data
	req := constants.GetDalleEndpoint("editImages")
	// Initialise the request body
	body := EditImagesPrompt{
		Prompt:   prompt,
		N:        n,
		Size:     size,
		UserName: userId,
	}
	// Set the body
	req.Body = body
	// Return an EditImagesRequest struct, this is used to add the image and mask to the request
	return &EditImagesRequest{
		req: req,
	}
}

// AddImageFromFileSystem adds an image to the EditImagesRequest
func (e *EditImagesRequest) AddImageFromFileSystem(image *os.File) *EditImagesRequest {
	im, err := io.ReadAll(image)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	e.Image = im
	return e
}

// AddMaskFromFileSystem adds a mask to the EditImagesRequest
func (e *EditImagesRequest) AddMaskFromFileSystem(mask *os.File) *EditImagesRequest {
	im, err := io.ReadAll(mask)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	e.Mask = im
	return e
}

// AddImage adds bytes to EditImagesRequest, this is used to add the image to the request
func (e *EditImagesRequest) AddImage(image []byte) *EditImagesRequest {
	e.Image = image
	return e
}

// AddMask adds bytes to EditImagesRequest, this is used to add the mask to the request
func (e *EditImagesRequest) AddMask(mask []byte) *EditImagesRequest {
	e.Mask = mask
	return e
}

// Do fires the request, returns an EditImageResponse
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
