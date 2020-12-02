package model

type VirtualMachineDisk struct {
	Size int `json:"size,omitempty"`
}

type VirtualMachineNetworkInterfaceIPAddress struct {
	IP      string        `json:"ip,omitempty"`
	Type    IPNetworkType `json:"type,omitempty"`
	RDNS    string        `json:"rdns,omitempty"`
	Netmask string        `json:"netmask,omitempty"`
	Gateway string        `json:"gateway,omitempty"`
	Primary bool          `json:"primary,omitempty"`
}

type VirtualMachineNetworkInterface struct {
	ID                  string                                    `json:"id,omitempty"`
	IPAddresses         []VirtualMachineNetworkInterfaceIPAddress `json:"ipAddresses,omitempty"`
	MAC                 string                                    `json:"mac,omitempty"`
	IPV6TransferNetwork string                                    `json:"ipv6TransferNetwork,omitempty"`
	IPV6RoutedNetwork   string                                    `json:"ipv6RoutedNetwork,omitempty"`
	IPV4PrimaryAddress  string                                    `json:"ipv4PrimaryAddress,omitempty"`
	IPV6PrimaryAddress  string                                    `json:"ipv6PrimaryAddress,omitempty"`
	NetworkID           string                                    `json:"networkId,omitempty"`
	NetworkType         string                                    `json:"networkType,omitempty"`
	NetworkBandwidth    int                                       `json:"networkBandwidth,omitempty"`
	Driver              string                                    `json:"driver,omitempty"`
}

type VirtualMachineManagementType string

const (
	// VirtualMachineManagementTypeNone is defined for a virtual machine that is unmanaged
	VirtualMachineManagementTypeNone VirtualMachineManagementType = "none"

	// VirtualMachineManagementTypeManaged is defined for a virtual machine that is (platform-) managed
	VirtualMachineManagementTypeManaged VirtualMachineManagementType = "managed"

	// VirtualMachineManagementTypeIndividual is defined for a virtual machine that is individually managed
	VirtualMachineManagementTypeIndividual VirtualMachineManagementType = "individual"
)

type VirtualMachineObject struct {
	ID                   string                           `json:"id,omitempty"`
	AccountID            string                           `json:"accountId,omitempty"`
	Name                 string                           `json:"name,omitempty"`
	Description          string                           `json:"description,omitempty"`
	ProductCode          string                           `json:"productCode,omitempty"`
	Memory               int                              `json:"memory,omitempty"`
	CPUNumber            int                              `json:"cpuNumber,omitempty"`
	Architecture         string                           `json:"architecture,omitempty"`
	Status               string                           `json:"status,omitempty"`
	IPAddress            string                           `json:"ipAddress,omitempty"`
	HostName             string                           `json:"hostName,omitempty"`
	RDNS                 string                           `json:"rdns,omitempty"`
	Power                string                           `json:"power,omitempty"`
	Rescue               string                           `json:"rescue,omitempty"`
	DiskControllerType   string                           `json:"diskControllerType,omitempty"`
	TimeZone             string                           `json:"timeZone,omitempty"`
	OSOptimization       string                           `json:"osOptimization,omitempty"`
	Disks                []VirtualMachineDisk             `json:"disks,omitempty"`
	BlockedPorts         []int                            `json:"blockedPorts,omitempty"`
	NetworkInterfaces    []VirtualMachineNetworkInterface `json:"networkInterfaces,omitempty"`
	ManagementType       VirtualMachineManagementType     `json:"managementType,omitempty"`
	ResourcePoolID       string                           `json:"resourcePoolId,omitempty"`
	VirtualMachineHostID string                           `json:"virtualMachineHostId,omitempty"`
	PaidUntil            string                           `json:"paidUntil,omitempty"`
	RenewOn              string                           `json:"renewOn,omitempty"`
	DeletionScheduledFor string                           `json:"deletionScheduledFor,omitempty"`
	RestorableUnit       string                           `json:"restorableUnit,omitempty"`
	AddDate              string                           `json:"addDate,omitempty"`
	LastChangeDate       string                           `json:"lastChangeDate,omitempty"`
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
