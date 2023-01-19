package gopenai

import (
	"encoding/json"
	"fmt"
	"github.com/LegalForceLawRAPC/gopenai/constants"
	"github.com/LegalForceLawRAPC/gopenai/errors"
	"io"
	"net/http"
)

func (c *Client) Do(data RequestData, i interface{}) *errors.HttpError {
	req, err := http.NewRequest(data.method, fmt.Sprintf("%s/%s", constants.BaseURL, data.endpoint), nil)
	if err != nil {
		return &errors.ErrDefault
	}
	req.Header.Add("Authorization", c.getBearerToken())
	res, err := c.client.Do(req)
	if err != nil {
		return errors.HandleHttpError(err, res.StatusCode)
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.HandleHttpError(err, res.StatusCode)
	}
	err = json.Unmarshal(bytes, &i)
	if err != nil {
		return errors.HandleHttpError(err, res.StatusCode)
	}
	return nil
}

func (c *Client) getBearerToken() string {
	return fmt.Sprintf("Bearer %s", c.basicAuth.apiKey)
}
