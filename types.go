package platform

type IPNetworkType string

const (
	IPNetworkTypeIPv4 IPNetworkType = "v4"
	IPNetworkTypeIPv6 IPNetworkType = "v6"
)

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

type VirtualMachineDisk struct {
	Size int `json:"size,omitempty"`
}

type VirtualMachineNetworkInterfaceIPAddress struct {
	Gateway string        `json:"gateway,omitempty"`
	IP      string        `json:"ip,omitempty"`
	Netmask string        `json:"netmask,omitempty"`
	Primary bool          `json:"primary,omitempty"`
	RDNS    string        `json:"rdns,omitempty"`
	Type    IPNetworkType `json:"type,omitempty"`
}

type VirtualMachineNetworkInterface struct {
	IPV4PrimaryAddress  string `json:"ipv4PrimaryAddress,omitempty"`
	IPV6PrimaryAddress  string `json:"ipv6PrimaryAddress,omitempty"`
	IPV6RoutedNetwork   string `json:"ipv6RoutedNetwork,omitempty"`
	IPV6TransferNetwork string `json:"ipv6TransferNetwork,omitempty"`
	MAC                 string `json:"mac,omitempty"`

	IPAddresses []VirtualMachineNetworkInterfaceIPAddress `json:"ipAddresses,omitempty"`
}

type VirtualMachineObject struct {
	AccountID            string                           `json:"accountId,omitempty"`
	AddDate              string                           `json:"addDate,omitempty"`
	Architecture         string                           `json:"architecture,omitempty"`
	BlockedPorts         []int                            `json:"blockedPorts,omitempty"`
	CPUNumber            int                              `json:"cpuNumber,omitempty"`
	DeletionScheduledFor string                           `json:"deletionScheduledFor,omitempty"`
	Description          string                           `json:"description,omitempty"`
	DiskControllerType   string                           `json:"diskControllerType,omitempty"`
	Disks                []VirtualMachineDisk             `json:"disks,omitempty"`
	HostName             string                           `json:"hostName,omitempty"`
	ID                   string                           `json:"id,omitempty"`
	IPAddress            string                           `json:"ipAddress,omitempty"`
	LastChangeDate       string                           `json:"lastChangeDate,omitempty"`
	ManagementType       string                           `json:"managementType,omitempty"`
	Memory               int                              `json:"memory,omitempty"`
	Name                 string                           `json:"name,omitempty"`
	NetworkInterfaces    []VirtualMachineNetworkInterface `json:"networkInterfaces,omitempty"`
	OSOptimization       string                           `json:"osOptimization,omitempty"`
	PaidUntil            string                           `json:"paidUntil,omitempty"`
	Power                string                           `json:"power,omitempty"`
	ProductCode          string                           `json:"productCode,omitempty"`
	RDNS                 string                           `json:"rdns,omitempty"`
	RenewOn              string                           `json:"renewOn,omitempty"`
	Rescue               string                           `json:"rescue,omitempty"`
	ResourcePoolID       string                           `json:"resourcePoolId,omitempty"`
	RestorableUnit       string                           `json:"restorableUnit,omitempty"`
	Status               string                           `json:"status,omitempty"`
	TimeZone             string                           `json:"timeZone,omitempty"`
	VirtualMachineHostID string                           `json:"virtualMachineHostId,omitempty"`
}

type FindVirtualMachinesResult struct {
	GenericResultMetadata

	Data []VirtualMachineObject `json:"data,omitempty"`
}

type VirtualMachineResult struct {
	Metadata Metadata                  `json:"metadata,omitempty"`
	Response FindVirtualMachinesResult `json:"response,omitempty"`

	Status   string
	Warnings []string
}

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
