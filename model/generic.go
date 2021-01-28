package model

type GenericResult struct {
	GenericResultMetadata

	Data []interface{} `json:"data,omitempty"`
}
