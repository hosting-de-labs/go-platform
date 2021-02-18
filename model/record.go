package model

type RecordObject struct {
	ID               string `json:"id,omitempty"`
	AccountID        string `json:"accountId,omitempty"`
	ZoneConfigID     string `json:"zoneConfigId,omitempty"`
	RecordTemplateID string `json:"recordTemplateId,omitempty"`
	Name             string `json:"name,omitempty"`
	Type             string `json:"type,omitempty"`
	Content          string `json:"content,omitempty"`
	TTL              int    `json:"ttl,omitempty"`
	Priority         int    `json:"priority,omitempty"`
	Comments         string `json:"comments,omitempty"`
	AddDate          string `json:"addDate,omitempty"`
	LastChangeDate   string `json:"lastChangeDate,omitempty"`
}
