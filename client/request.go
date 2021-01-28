package client

import (
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
	reqBody := &GenericRequestBody{}
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

	requestURL := fmt.Sprintf("%s/%s/v1/json/%s", c.url, endpoint, rpcMethod)

	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		Post(requestURL)

	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}

	return resp.RawResponse, nil
}
