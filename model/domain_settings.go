package model

type DomainSettingsObject struct {
	AddDate                       string `json:"addDate,omitempty"`
	BundleID                      string `json:"bundleId,omitempty"`
	DomainName                    string `json:"domainName,omitempty"`
	DomainNameUnicode             string `json:"domainNameUnicode,omitempty"`
	ExchangeMailboxQuota          int    `json:"exchangeMailboxQuota,omitempty"`
	ExchangeStorageQuota          int    `json:"exchangeStorageQuota,omitempty"`
	ExchangeStorageQuotaAllocated int    `json:"exchangeStorageQuotaAllocated,omitempty"`
	LastChangeDate                string `json:"lastChangeDate,omitempty"`
	MailboxQuota                  int    `json:"mailboxQuota,omitempty"`
	StorageQuota                  int    `json:"storageQuota,omitempty"`
	StorageQuotaAllocated         int    `json:"storageQuotaAllocated,omitempty"`
}

type FindDomainSettingsResult struct {
	GenericResultMetadata

	Data []DomainSettingsObject `json:"data,omitempty"`
}

type DomainSettingsResult struct {
	Metadata Metadata                 `json:"metadata,omitempty"`
	Response FindDomainSettingsResult `json:"response,omitempty"`

	Status   string
	Warnings []string
}
