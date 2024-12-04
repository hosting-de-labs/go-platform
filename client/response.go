package client

import (
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

func ParseFindResponse(out *[]interface{}, T interface{}, resp *resty.Response) (int, int, error) {
	var currentPageBody FindResponse
	err := json.Unmarshal(resp.Body(), &currentPageBody)
	if err != nil {
		return -1, -1, fmt.Errorf("parse: unmarshal failed: %s", err)
	}

	if len(currentPageBody.Errors) > 0 {
		return -1, -1, fmt.Errorf("server reported error: %s", currentPageBody.Errors[0].Text)
	}

	for _, data := range currentPageBody.Response.Data {
		err := DecodeFindData(data, T)
		if err != nil {
			return -1, -1, err
		}

		*out = append(*out, T)
	}

	return len(currentPageBody.Response.Data), currentPageBody.Response.TotalPages, nil
}

func DecodeFindData(data interface{}, T interface{}) error {
	err := mapstructure.Decode(data, &T)
	if err != nil {
		return fmt.Errorf("parse: decode failed: %s", err)
	}
	return nil
}
