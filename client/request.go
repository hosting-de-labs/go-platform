package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/hosting-de-labs/go-platform/model"
)

const (
	IterateDefaultLimit = 250
)

type GenericRequestBody struct {
	AuthToken string         `json:"authToken,omitempty"`
	Limit     int            `json:"limit,omitempty"`
	Page      int            `json:"page,omitempty"`
	Filter    *RequestFilter `json:"filter,omitempty"`
}

func (c *ApiClient) Iterate(data []interface{}, endpoint string, rpcMethod string, filter *RequestFilter, page int) (int, error) {
	totalObjects := 0
	for i := 0; ; i++ {
		currentPage := 1 + i
		if page >= 1 {
			currentPage = page + i
		}

		resp, err := c.Request(endpoint, rpcMethod, filter, IterateDefaultLimit, currentPage)
		if err != nil {
			return 0, fmt.Errorf("iterate failed: %s", err)
		}

		var currentPageBody model.GenericResult
		err = json.Unmarshal(resp.Body(), &currentPageBody)
		if err != nil {
			return 0, err
		}

		totalObjects += len(currentPageBody.Data)

		data = append(data, currentPageBody.Data...)
		if currentPage == currentPageBody.TotalPages {
			break
		}
	}

	return totalObjects, nil
}

// Request fires a generic request.
func (c *ApiClient) Request(endpoint string, rpcMethod string, filter *RequestFilter, limit int, page int) (*resty.Response, error) {
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

	return resp, nil
}

// runRequest
// Deprecated: Use Request() method insted
func (c *ApiClient) runRequest(endpoint string, rpcMethod string, filter *RequestFilter, limit int, page int) (*http.Response, error) {
	resp, err := c.Request(endpoint, rpcMethod, filter, limit, page)
	if err != nil {
		return nil, err
	}

	return resp.RawResponse, nil
}
