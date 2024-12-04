package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
)

type Metadata struct {
	ClientTransactionID string `json:"clientTransactionId,omitempty"`
	ServerTransactionID string `json:"serverTransactionId,omitempty"`
}

type EmptyResponse struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Status   string   `json:"status,omitempty"`
	Errors   []Error  `json:"errors,omitempty"`
	Warnings []Error  `json:"warnings,omitempty"`
}

type FindResponseMetadata struct {
	Limit        int    `json:"limit,omitempty"`
	Page         int    `json:"page,omitempty"`
	TotalEntries int    `json:"totalEntries,omitempty"`
	TotalPages   int    `json:"totalPages,omitempty"`
	Type         string `json:"type,omitempty"`
}

type FindResponse struct {
	EmptyResponse
	Response struct {
		FindResponseMetadata

		Data []map[string]interface{} `json:"data,omitempty"`
	} `json:"response,omitempty"`
}

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

		findResp := &FindResponse{}
		err := c.Request(endpoint, rpcMethod, reqBody, findResp, &FindResponse{})
		if err != nil {
			return -1, err
		}

		if len(findResp.Errors) > 0 {
			return -1, fmt.Errorf("server reported error: %s", findResp.Errors[0].Text)
		}

		err = iterateFindResponse(findResp, data, T)
		if err != nil {
			return -1, err
		}

		totalObjects += len(findResp.Response.Data)
		if page >= findResp.Response.TotalPages {
			break
		}
	}

	return totalObjects, nil
}

func (c *ApiClient) Request(endpoint string, rpcMethod string, body interface{}, out interface{}, T interface{}) error {
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

// RawRequest fires a generic request that carries data to the server
func (c *ApiClient) RawRequest(endpoint string, rpcMethod string, body interface{}) (*resty.Response, error) {
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

func decodeFindData(data interface{}, T interface{}) error {
	err := mapstructure.Decode(data, &T)
	if err != nil {
		return fmt.Errorf("parse: decode failed: %s", err)
	}
	return nil
}

func iterateFindResponse(findResp *FindResponse, out *[]interface{}, T interface{}) error {
	for _, data := range findResp.Response.Data {
		err := decodeFindData(data, T)
		if err != nil {
			return err
		}

		*out = append(*out, T)
	}

	return nil
}
