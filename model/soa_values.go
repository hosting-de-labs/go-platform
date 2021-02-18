package model

type SoaValuesObject struct {
	Refresh     int `json:"refresh,omitempty"`
	Retry       int `json:"retry,omitempty"`
	Expire      int `json:"expire,omitempty"`
	TTL         int `json:"ttl,omitempty"`
	NegativeTTL int `json:"negativeTtl,omitempty"`
}
