package model

type ZoneConfigObject struct {
	ID                    string               `json:"id,omitempty"`
	AccountID             string               `json:"accountId,omitempty"`
	DNSServerGroupID      string               `json:"dnsServerGroupId,omitempty"`
	Status                string               `json:"status,omitempty"`
	Name                  string               `json:"name,omitempty"`
	NameUnicode           string               `json:"nameUnicode,omitempty"`
	MasterIP              string               `json:"masterIp,omitempty"`
	Type                  string               `json:"type,omitempty"`
	EmailAddress          string               `json:"emailAddress,omitempty"`
	ZoneTransferWhitelist []string             `json:"zoneTransferWhitelist,omitempty"`
	SoaValues             SoaValuesObject      `json:"soaValues,omitempty"`
	TemplateValues        TemplateValuesObject `json:"templateValues,omitempty"`
	DNSSecMode            string               `json:"dnsSecMode,omitempty"`
	Restrictions          []RestrictionObject  `json:"restrictions,omitempty"`
	LastChangeDate        string               `json:"lastChangeDate,omitempty"`
	AddDate               string               `json:"addDate,omitempty"`
}
