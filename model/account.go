package model

type BillingSettings struct {
	CreditLimit                          float64                `json:"creditLimit,omitempty"`
	DataProcessingAgreementSignedOn      string                 `json:"dataProcessingAgreementSignedOn,omitempty"`
	DataProcessingAgreementSignedVersion int                    `json:"dataProcessingAgreementSignedVersion,omitempty"`
	PaymentType                          string                 `json:"paymentType,omitempty"`
	PostpaidType                         string                 `json:"postpaidType,omitempty"`
	PurchaseMarkup                       string                 `json:"purchaseMarkup,omitempty"`
	RenewalsReportSettings               RenewalsReportSettings `json:"renewalsReportSettings,omitempty"`
	Type                                 BillingSettingsType    `json:"type,omitempty"`
}

type RenewalsReportSettings struct {
	Inverval                      int `json:"creditLimit,omitempty"`
	Horizon                       int `json:"horizon,omitempty"`
	ReceiveSufficientFundsReports int `json:"receiveSufficientFundsReports,omitempty"`
}

type DnsSettings struct {
	DefaultHostmasterEmailAddress string          `json:"defaultHostmasterEmailAddress,omitempty"`
	DefaultTemplateId             string          `json:"defaultTemplateId,omitempty"`
	Type                          DnsSettingsType `json:"type,omitempty"`
}

type DomainSettings struct {
	DefaultContactAdminId string             `json:"defaultContactAdminId,omitempty"`
	DefaultContactOwnerId string             `json:"defaultContactOwnerId,omitempty"`
	DefaultContactTechId  string             `json:"defaultContactTechId,omitempty"`
	DefaultContactZoneId  string             `json:"defaultContactZoneId,omitempty"`
	MonthlyPaymentEnabled bool               `json:"monthlyPaymentEnabled,omitempty"`
	Type                  DomainSettingsType `json:"type,omitempty"`
}

type ResellerBillingSettings struct {
	ActivePaymentMethods                       []string `json:"activePaymentMethods,omitempty"`
	AvailablePaymentMethods                    []string `json:"availablePaymentMethods,omitempty"`
	DefaultCreditLimitForCommercialSubaccounts int      `json:"defaultCreditLimitForCommercialSubaccounts,omitempty"`
	DefaultCreditLimitForEndUserSubaccounts    int      `json:"defaultCreditLimitForEndUserSubaccounts,omitempty"`
	EndUserCancellationPolicy                  string   `json:"endUserCancellationPolicy,omitempty"`
	RetailMarkup                               int      `json:"retailMarkup,omitempty"`
	ShowEndUserCancellationPolicyAboveAmount   int      `json:"showEndUserCancellationPolicyAboveAmount,omitempty"`
}

type DefaultMxRecord struct {
	Content  string `json:"content,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type ResellerDefaults struct {
	DefaultDnsNameserverSet           string            `json:"defaultDnsNameserverSet,omitempty"`
	DefaultDnsTemplate                string            `json:"defaultDnsTemplate,omitempty"`
	DefaultMailServer                 string            `json:"defaultMailServer,omitempty"`
	DefaultMxRecords                  []DefaultMxRecord `json:"defaultMxRecords,omitempty"`
	DefaultSubaccountRightsTemplateId string            `json:"defaultSubaccountRightsTemplateId,omitempty"`
}

type CompanyBankAccount struct {
	AccountNumber string `json:"accountNumber,omitempty"`
	BankName      string `json:"bankName,omitempty"`
	BIC           string `json:"bic,omitempty"`
}

type ResellerWhiteLabelSettings struct {
	Address                             string             `json:"address,omitempty"`
	AlternativeEmailSender              string             `json:"alternativeEmailSender,omitempty"`
	CompanyBankAccount                  CompanyBankAccount `json:"companyBankAccount,omitempty"`
	ConfirmationUrl                     string             `json:"confirmationUrl,omitempty"`
	DomainRegistrationTermsUrl          string             `json:"domainRegistrationTermsUrl,omitempty"`
	EmailAddress                        string             `json:"emailAddress,omitempty"`
	EmailAddressForActionRequiredEmails string             `json:"emailAddressForActionRequiredEmails,omitempty"`
	EmailAddressForInfoOnlyEmails       string             `json:"emailAddressForInfoOnlyEmails,omitempty"`
	Name                                string             `json:"name,omitempty"`
	PhoneNumber                         string             `json:"phoneNumber,omitempty"`
	Website                             string             `json:"website,omitempty"`
}

type MessageSettings struct {
	DefaultHtmlFrame                   string              `json:"defaultHtmlFrame,omitempty"`
	DefaultPlainTextFrame              string              `json:"defaultPlainTextFrame,omitempty"`
	EventDeliveryMethods               []string            `json:"eventDeliveryMethods,omitempty"`
	EventDocumentFormat                string              `json:"eventDocumentFormat,omitempty"`
	EventHumanReadableEmailAddresses   []string            `json:"eventHumanReadableEmailAddresses,omitempty"`
	EventMachineReadableEmailAddresses []string            `json:"eventMachineReadableEmailAddresses,omitempty"`
	EventWebhookUrl                    string              `json:"eventWebhookUrl,omitempty"`
	MessageTemplateBehavior            string              `json:"messageTemplateBehavior,omitempty"`
	ReceiveSubaccountEvents            bool                `json:"receiveSubaccountEvents,omitempty"`
	Type                               MessageSettingsType `json:"type,omitempty"`
}

type AccountObject struct {
	ID                           string                     `json:"id,omitempty"`
	ParentAccountId              string                     `json:"parentAccountId,omitempty"`
	Type                         AccountType                `json:"type,omitempty"`
	RightsTemplateId             string                     `json:"rightsTemplateId,omitempty"`
	IsCommercialCustomer         bool                       `json:"isCommercialCustomer,omitempty"`
	CustomerNumber               string                     `json:"customerNumber,omitempty"`
	ContactName                  string                     `json:"contactName,omitempty"`
	ContactTitle                 string                     `json:"contactTitle,omitempty"`
	Name                         string                     `json:"name,omitempty"`
	Title                        string                     `json:"title,omitempty"`
	Addition                     string                     `json:"addition,omitempty"`
	Street                       string                     `json:"street,omitempty"`
	ZipCode                      string                     `json:"zipCode,omitempty"`
	City                         string                     `json:"city,omitempty"`
	PhoneNumber                  string                     `json:"phoneNumber,omitempty"`
	FaxNumber                    string                     `json:"faxNumber,omitempty"`
	EmailAddress                 string                     `json:"emailAddress,omitempty"`
	EmailAddressInvoice          string                     `json:"emailAddressInvoice,omitempty"`
	EmailAddressNewsletter       string                     `json:"emailAddressNewsletter,omitempty"`
	EmailAddressTechnicalContact string                     `json:"emailAddressTechnicalContact,omitempty"`
	Country                      string                     `json:"country,omitempty"`
	Language                     string                     `json:"language,omitempty"`
	AddDate                      string                     `json:"addDate,omitempty"`
	LastChangeDate               string                     `json:"lastChangeDate,omitempty"`
	BillingSettings              BillingSettings            `json:"billingSettings,omitempty"`
	Comments                     string                     `json:"comments,omitempty"`
	Enabled                      bool                       `json:"enabled,omitempty"`
	DnsSettings                  DnsSettings                `json:"dnsSettings,omitempty"`
	DomainSettings               DomainSettings             `json:"domainSettings,omitempty"`
	MessageSettings              MessageSettings            `json:"messageSettings,omitempty"`
	ResellerBillingSettings      ResellerBillingSettings    `json:"resellerBillingSettings,omitempty"`
	ResellerDefaults             ResellerDefaults           `json:"resellerDefaults,omitempty"`
	ResellerWhiteLabelSettings   ResellerWhiteLabelSettings `json:"resellerWhiteLabelSettings,omitempty"`
	Rights                       []string                   `json:"rights,omitempty"`
	SalesTaxID                   string                     `json:"salesTaxID,omitempty"`
	AllowedAdvertisingMethods    []string                   `json:"allowedAdvertisingMethods,omitempty"`
	ReferralCode                 string                     `json:"referralCode,omitempty"`
	ReferringAccountId           string                     `json:"referringAccountId,omitempty"`
	Verifications                []string                   `json:"verifications,omitempty"`
}

type DnsSettingsType string

const (
	DnsSettingsTypeSubaccountDnsSettings DnsSettingsType = "SubaccountDnsSettings"
)

type AccountType string

const (
	AccountTypeSubaccount AccountType = "Subaccount"
)

type BillingSettingsType string

const (
	BillingSettingsTypeSubaccountBillingSettings BillingSettingsType = "SubaccountBillingSettings"
)

type DomainSettingsType string

const (
	DomainSettingsTypeSubaccountDomainSettings DomainSettingsType = "SubaccountDomainSettings"
)

type MessageSettingsType string

const (
	MessageSettingsTypeSubaccountMessageSettings MessageSettingsType = "SubaccountMessageSettings"
)
