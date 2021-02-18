package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
)

const (
	IterateDefaultLimit = 250
)

type GetRequest struct {
	Limit  int            `json:"limit,omitempty"`
	Page   int            `json:"page,omitempty"`
	Filter *RequestFilter `json:"filter,omitempty"`
}

func (c *ApiClient) Iterate(data *[]interface{}, T interface{}, endpoint string, rpcMethod string, filter *RequestFilter, page int) (int, error) {
	totalObjects := 0
	for i := 0; ; i++ {
		currentPage := 1 + i
		if page >= 1 {
			currentPage = page + i
		}

		resp, err := c.Get(endpoint, rpcMethod, filter, IterateDefaultLimit, currentPage)
		if err != nil {
			return 0, fmt.Errorf("iterate failed: %s", err)
		}

		var currentPageBody DataResponse
		err = json.Unmarshal(resp.Body(), &currentPageBody)
		if err != nil {
			return 0, fmt.Errorf("iterate failed: %s", err)
		}

		if len(currentPageBody.Errors) > 0 {
			return 0, fmt.Errorf("iterate: %s", currentPageBody.Errors[0].Text)
		}

		totalObjects += len(currentPageBody.Response.Data)

		for _, r := range currentPageBody.Response.Data {
			err := mapstructure.Decode(r, &T)
			if err != nil {
				panic(err)
			}

			*data = append(*data, T)
		}

		if currentPage >= currentPageBody.Response.TotalPages {
			break
		}
	}

	return totalObjects, nil
}

// Get fires a generic request that will not push any further data to the api.
func (c *ApiClient) Get(endpoint string, rpcMethod string, filter *RequestFilter, limit int, page int) (*resty.Response, error) {
	reqBody := &GetRequest{}

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

	base64Token := base64.StdEncoding.EncodeToString([]byte(c.token))
	resp, err := c.client.R().
		SetHeader("Authorization", "Bearer "+base64Token).
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		Post(requestURL)

	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}

	return resp, nil
}

// Update fires a generic request that carries data to the server
func (c *ApiClient) Update(endpoint string, rpcMethod string, body interface{}) (*resty.Response, error) {
	requestURL := fmt.Sprintf("%s/%s/v1/json/%s", c.url, endpoint, rpcMethod)

	base64Token := base64.StdEncoding.EncodeToString([]byte(c.token))
	resp, err := c.client.R().
		SetHeader("Authorization", "Bearer "+base64Token).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(requestURL)

	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}

	return resp, nil
}
