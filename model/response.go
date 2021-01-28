package model

type Metadata struct {
	ClientTransactionID string `json:"clientTransactionId,omitempty"`
	ServerTransactionID string `json:"serverTransactionId,omitempty"`
}

type ResponseMetadata struct {
	Limit        int    `json:"limit,omitempty"`
	Page         int    `json:"page,omitempty"`
	TotalEntries int    `json:"totalEntries,omitempty"`
	TotalPages   int    `json:"totalPages,omitempty"`
	Type         string `json:"type,omitempty"`
}

type Response struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Response struct {
		ResponseMetadata

		Data []map[string]interface{} `json:"data,omitempty"`
	} `json:"response,omitempty"`

	Status string
	Errors []struct {
		Code          int
		ContextObject string
		ContextPath   string
		Details       map[string]string
		Text          string
		Value         string
	}
}
