package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) MakeHTTPRequest(url interface{}, method string, endpoint string, reqBody []byte) ([]byte, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, fmt.Sprintf("%v", url)+endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-api-key", c.secretKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode > http.StatusCreated {
		var errorMessage struct {
			Message string `json:"message"`
		}

		return nil, errors.New(errorMessage.Message)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
