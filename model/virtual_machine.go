package model

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
	ManagementType       VirtualMachineManagementType     `json:"managementType,omitempty"`
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
