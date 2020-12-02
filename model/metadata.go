package model

type Metadata struct {
	ClientTransactionID string `json:"clientTransactionId,omitempty"`
	ServerTransactionID string `json:"serverTransactionId,omitempty"`
}

type GenericResultMetadata struct {
	Limit        int    `json:"limit,omitempty"`
	Page         int    `json:"page,omitempty"`
	TotalEntries int    `json:"totalEntries,omitempty"`
	TotalPages   int    `json:"totalPages,omitempty"`
	Type         string `json:"type,omitempty"`
}
