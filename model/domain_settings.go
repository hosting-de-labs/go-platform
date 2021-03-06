package model

type DomainSettingsObject struct {
	BundleID                      string `json:"bundleId,omitempty"`
	DomainName                    string `json:"domainName,omitempty"`
	DomainNameUnicode             string `json:"domainNameUnicode,omitempty"`
	StorageQuota                  int    `json:"storageQuota,omitempty"`
	StorageQuotaAllocated         int    `json:"storageQuotaAllocated,omitempty"`
	MailboxQuota                  int    `json:"mailboxQuota,omitempty"`
	ExchangeMailboxQuota          int    `json:"exchangeMailboxQuota,omitempty"`
	ExchangeStorageQuota          int    `json:"exchangeStorageQuota,omitempty"`
	ExchangeStorageQuotaAllocated int    `json:"exchangeStorageQuotaAllocated,omitempty"`
	AddDate                       string `json:"addDate,omitempty"`
	LastChangeDate                string `json:"lastChangeDate,omitempty"`
}
