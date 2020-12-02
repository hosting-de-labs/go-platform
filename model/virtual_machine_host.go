package model

type VirtualMachineHostObject struct {
	ID                string `json:"id,omitempty"`
	AccountID         string `json:"accountId,omitempty"`
	PoolID            string `json:"poolId,omitempty"`
	ResourceType      string `json:"resourceType,omitempty"`
	HostName          string `json:"hostName,omitempty"`
	Priority          int    `json:"priority,omitempty"`
	Ignore            bool   `json:"ignore,omitempty"`
	AddDate           string `json:"addDate,omitempty"`
	LastChangeDate    string `json:"lastChangeDate,omitempty"`
	PhysicalMachineID string `json:"physicalMachineId,omitempty"`
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
