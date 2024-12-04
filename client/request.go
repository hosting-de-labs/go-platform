package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	IterateDefaultLimit = 250
)

type findRequest struct {
	Limit  int            `json:"limit,omitempty"`
	Page   int            `json:"page,omitempty"`
	Filter *RequestFilter `json:"filter,omitempty"`
}

func (c *ApiClient) Find(data *[]interface{}, T interface{}, endpoint string, rpcMethod string, filter *RequestFilter, limit int) (int, error) {
	totalObjects := 0
	for page := 1; ; page++ {
		reqBody := &findRequest{}

		if filter != nil {
			reqBody.Filter = filter
		}

		reqBody.Limit = IterateDefaultLimit
		if limit > 0 {
			reqBody.Limit = limit
		}
		reqBody.Page = page

		resp, err := c.Request(endpoint, rpcMethod, reqBody)
		if err != nil {
			return -1, err
		}
		numObjects, totalPages, err := ParseFindResponse(data, T, resp)
		if err != nil {
			return -1, err
		}

		totalObjects += numObjects
		if page >= totalPages {
			break
		}
	}

	return totalObjects, nil
}

func (c *ApiClient) ParsedRequest(endpoint string, rpcMethod string, body interface{}, out interface{}, T interface{}) error {
	resp, err := c.request(endpoint, rpcMethod, body)
	if err != nil {
		return fmt.Errorf("request: %s", err)
	}

	err = json.Unmarshal(resp.Body(), &out)
	if err != nil {
		return fmt.Errorf("parse: unmarshal failed: %s", err)
	}

	return nil
}

// Request fires a generic request that carries data to the server
func (c *ApiClient) Request(endpoint string, rpcMethod string, body interface{}) (*resty.Response, error) {
	resp, err := c.request(endpoint, rpcMethod, body)
	if err != nil {
		return nil, fmt.Errorf("request: %s", err)
	}

	return resp, nil
}

func (c *ApiClient) request(endpoint string, rpcMethod string, body interface{}) (*resty.Response, error) {
	reqURL := fmt.Sprintf("%s/%s/v1/json/%s", c.url, endpoint, rpcMethod)
	base64Token := base64.StdEncoding.EncodeToString([]byte(c.token))
	resp, err := c.client.R().
		SetHeader("Authorization", "Bearer "+base64Token).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(reqURL)

	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}

	return resp, nil
}
