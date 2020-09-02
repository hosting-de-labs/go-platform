package model

type VirtualMachineHostObject struct {
	AccountID         string `json:"accountId,omitempty"`
	AddDate           string `json:"addDate,omitempty"`
	HostName          string `json:"hostName,omitempty"`
	ID                string `json:"id,omitempty"`
	Ignore            bool   `json:"ignore,omitempty"`
	LastChangeDate    string `json:"lastChangeDate,omitempty"`
	PhysicalMachineID string `json:"physicalMachineId,omitempty"`
	PoolID            string `json:"poolId,omitempty"`
	Priority          int    `json:"priority,omitempty"`
	Type              string `json:"type,omitempty"`
}

type FindVirtualMachineHostsResult struct {
	GenericResultMetadata

	Data []VirtualMachineHostObject `json:"data,omitempty"`
}

type VirtualMachineHostsResult struct {
	Metadata Metadata                      `json:"metadata,omitempty"`
	Response FindVirtualMachineHostsResult `json:"response,omitempty"`

	Status   string   `json:"status,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
