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

func (d *Dalle) Do(data constants.RequestData, i interface{}) *errors.HttpError {
	jsonData, err := json.Marshal(data.Body)
	if err != nil {
		return &errors.ErrDefault
	}
	req, err := http.NewRequest(data.Method, fmt.Sprintf("%s/%s", constants.BaseURL, data.Endpoint), bytes2.NewBuffer(jsonData))
	if err != nil {
		return &errors.ErrDefault
	}
	req.Header.Add("Authorization", constants.GetToken())
	if data.ContentType != "" {
		req.Header.Add("Content-Type", data.ContentType)
	} else {
		req.Header.Add("Content-Type", "application/json")
	}
	res, err := d.Client.Do(req)
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
