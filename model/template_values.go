package model

type TemplateValuesObject struct {
	ID             string `json:"id,omitempty"`
	AccountID      string `json:"accountId,omitempty"`
	Name           string `json:"name,omitempty"`
	EmailAddress   string `json:"emailAddress,omitempty"`
	LastChangeDate string `json:"lastChangeDate,omitempty"`
	AddDate        string `json:"addDate,omitempty"`
}
