package client

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

	Status string  `json:"status,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}
