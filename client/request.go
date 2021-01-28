package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GenericRequestBody struct {
	AuthToken string         `json:"authToken,omitempty"`
	Limit     int            `json:"limit,omitempty"`
	Page      int            `json:"page,omitempty"`
	Filter    *RequestFilter `json:"filter,omitempty"`
}

func (c *ApiClient) runRequest(endpoint string, rpcMethod string, filter *RequestFilter, limit int, page int) (*http.Response, error) {
	reqBody := new(GenericRequestBody)
	reqBody.AuthToken = c.token

	if filter != nil {
		reqBody.Filter = filter
	}

	reqBody.Limit = c.limit
	if limit > 0 {
		reqBody.Limit = limit
	}

	reqBody.Page = 1
	if page > 0 {
		reqBody.Page = page
	}

	requestBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/v1/json/%s", c.url, endpoint, rpcMethod), bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
