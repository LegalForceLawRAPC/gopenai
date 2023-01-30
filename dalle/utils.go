package dalle

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/LegalForceLawRAPC/gopenai/constants"
	"github.com/LegalForceLawRAPC/gopenai/errors"
)

// Do fires the request as needed
func (d *Dalle) Do(data constants.RequestData, i interface{}) *errors.HttpError {
	// Marshal the body into json
	jsonData, err := json.Marshal(data.Body)
	if err != nil {
		return &errors.ErrDefault
	}
	// Create the request
	req, err := http.NewRequest(data.Method, fmt.Sprintf("%s/%s", constants.BaseURL, data.Endpoint), bytes2.NewBuffer(jsonData))
	if err != nil {
		return &errors.ErrDefault
	}
	// Add the bearer token
	req.Header.Add("Authorization", constants.GetToken())
	// If ContentType is not specified, default to application/json
	if data.ContentType != "" {
		// Add the content type
		req.Header.Add("Content-Type", data.ContentType)
	} else {
		// Add the content type
		req.Header.Add("Content-Type", "application/json")
	}
	// Send the request
	res, err := d.Client.Do(req)
	if err != nil {
		return errors.HandleHttpError(err, res.StatusCode)
	}
	defer res.Body.Close()
	// Read the response
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		// If the status code is not 200, return an error
		return errors.HandleHttpError(err, res.StatusCode)
	}
	// Unmarshal the response into the interface
	err = json.Unmarshal(bytes, &i)
	if err != nil {
		// If the status code is not 200, return an error
		return errors.HandleHttpError(err, res.StatusCode)
	}

	return nil
}
