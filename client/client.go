package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiClient struct {
	url   string
	token string
	limit int
}

type GenericRequestBody struct {
	AuthToken string         `json:"authToken,omitempty"`
	Limit     int            `json:"limit,omitempty"`
	Page      int            `json:"page,omitempty"`
	Filter    *RequestFilter `json:"filter,omitempty"`
}

type SubFilterConnective string

type FilterRelation string

const (
	SubFilterConnectiveOr  SubFilterConnective = "OR"
	SubFilterConnectiveAnd SubFilterConnective = "AND"

	FilterRelationEqual        FilterRelation = "equal"
	FilterRelationUnequal      FilterRelation = "unequal"
	FilterRelationGreater      FilterRelation = "greater"
	FilterRelationLess         FilterRelation = "less"
	FilterRelationGreaterEqual FilterRelation = "greaterEqual"
	FilterRelationLessEqual    FilterRelation = "lessEqual"
)

type RequestFilter struct {
	Field    string          `json:"field,omitempty"`
	Value    string          `json:"value,omitempty"`
	Relation *FilterRelation `json:"relation,omitempty"`

	SubFilterConnective *SubFilterConnective `json:"subFilterConnective,omitempty"`
	SubFilter           *[]RequestFilter     `json:"subFilter,omitempty"`
}

func NewApiClient(url string, token string, limit int) *ApiClient {
	c := new(ApiClient)

	c.url = url
	c.token = token

	c.limit = 25
	if limit > 0 {
		c.limit = limit
	}

	return c
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
