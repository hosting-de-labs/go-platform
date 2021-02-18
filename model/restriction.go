package model

type RestrictionObject struct {
	ID              string `json:"id,omitempty"`
	Type            string `json:"type,omitempty"`
	Status          string `json:"status,omitempty"`
	AccountID       string `json:"accountId,omitempty"`
	User            string `json:"user,omitempty"`
	InternalComment string `json:"internalComment,omitempty"`
	Information     string `json:"information,omitempty"`
	AddDate         string `json:"addDate,omitempty"`
	LastChangeDate  string `json:"lastChangeDate,omitempty"`
}
