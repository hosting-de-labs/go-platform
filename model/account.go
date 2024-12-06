package model

type BillingSettings struct {
	CreditLimit                          float64                 `json:"creditLimit"`
	DataProcessingAgreementSignedOn      string                  `json:"dataProcessingAgreementSignedOn"`
	DataProcessingAgreementSignedVersion int                     `json:"dataProcessingAgreementSignedVersion"`
	PaymentType                          string                  `json:"paymentType"`
	PostpaidType                         string                  `json:"postpaidType"`
	PurchaseMarkup                       float64                 `json:"purchaseMarkup"`
	RenewalsReportSettings               *RenewalsReportSettings `json:"renewalsReportSettings"`
	Type                                 BillingSettingsType     `json:"type,omitempty"`
}

type RenewalsReportSettings struct {
	Interval                      int  `json:"interval"`
	Horizon                       int  `json:"horizon"`
	ReceiveSufficientFundsReports bool `json:"receiveSufficientFundsReports"`
}

type DnsSettings struct {
	DefaultHostmasterEmailAddress string          `json:"defaultHostmasterEmailAddress"`
	DefaultTemplateId             string          `json:"defaultTemplateId"`
	Type                          DnsSettingsType `json:"type,omitempty"`
}

type DomainSettings struct {
	DefaultContactAdminId string             `json:"defaultContactAdminId"`
	DefaultContactOwnerId string             `json:"defaultContactOwnerId"`
	DefaultContactTechId  string             `json:"defaultContactTechId"`
	DefaultContactZoneId  string             `json:"defaultContactZoneId"`
	MonthlyPaymentEnabled bool               `json:"monthlyPaymentEnabled"`
	Type                  DomainSettingsType `json:"type,omitempty"`
}

type ResellerBillingSettings struct {
	ActivePaymentMethods                       []string `json:"activePaymentMethods"`
	AvailablePaymentMethods                    []string `json:"availablePaymentMethods"`
	DefaultCreditLimitForCommercialSubaccounts int      `json:"defaultCreditLimitForCommercialSubaccounts"`
	DefaultCreditLimitForEndUserSubaccounts    int      `json:"defaultCreditLimitForEndUserSubaccounts"`
	EndUserCancellationPolicy                  string   `json:"endUserCancellationPolicy"`
	RetailMarkup                               int      `json:"retailMarkup"`
	ShowEndUserCancellationPolicyAboveAmount   int      `json:"showEndUserCancellationPolicyAboveAmount"`
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
	AccountNumber string `json:"accountNumber"`
	BankName      string `json:"bankName"`
	BIC           string `json:"bic"`
}

type ResellerWhiteLabelSettings struct {
	Address                             string             `json:"address"`
	AlternativeEmailSender              string             `json:"alternativeEmailSender,omitempty"`
	CompanyBankAccount                  CompanyBankAccount `json:"companyBankAccount,omitempty"`
	ConfirmationUrl                     string             `json:"confirmationUrl,omitempty"`
	DomainRegistrationTermsUrl          string             `json:"domainRegistrationTermsUrl,omitempty"`
	EmailAddress                        string             `json:"emailAddress"`
	EmailAddressForActionRequiredEmails string             `json:"emailAddressForActionRequiredEmails,omitempty"`
	EmailAddressForInfoOnlyEmails       string             `json:"emailAddressForInfoOnlyEmails,omitempty"`
	Name                                string             `json:"name"`
	PhoneNumber                         string             `json:"phoneNumber"`
	Website                             string             `json:"website,omitempty"`
}

type MessageSettings struct {
	DefaultHtmlFrame                   string              `json:"defaultHtmlFrame"`
	DefaultPlainTextFrame              string              `json:"defaultPlainTextFrame"`
	EventDeliveryMethods               []string            `json:"eventDeliveryMethods"`
	EventDocumentFormat                string              `json:"eventDocumentFormat"`
	EventHumanReadableEmailAddresses   []string            `json:"eventHumanReadableEmailAddresses"`
	EventMachineReadableEmailAddresses []string            `json:"eventMachineReadableEmailAddresses"`
	EventWebhookUrl                    string              `json:"eventWebhookUrl"`
	MessageTemplateBehavior            string              `json:"messageTemplateBehavior"`
	ReceiveSubaccountEvents            bool                `json:"receiveSubaccountEvents"`
	Type                               MessageSettingsType `json:"type"`
}

type AccountObject struct {
	ID                           string                      `json:"id,omitempty"`
	ParentAccountId              string                      `json:"parentAccountId,omitempty"`
	Type                         AccountType                 `json:"type,omitempty"`
	RightsTemplateId             string                      `json:"rightsTemplateId,omitempty"`
	IsCommercialCustomer         bool                        `json:"isCommercialCustomer"`
	CustomerNumber               string                      `json:"customerNumber,omitempty"`
	ContactName                  string                      `json:"contactName,omitempty"`
	ContactTitle                 string                      `json:"contactTitle,omitempty"`
	Name                         string                      `json:"name,omitempty"`
	Title                        string                      `json:"title,omitempty"`
	Addition                     string                      `json:"addition,omitempty"`
	Street                       string                      `json:"street,omitempty"`
	ZipCode                      string                      `json:"zipCode,omitempty"`
	City                         string                      `json:"city,omitempty"`
	PhoneNumber                  string                      `json:"phoneNumber,omitempty"`
	FaxNumber                    string                      `json:"faxNumber,omitempty"`
	EmailAddress                 string                      `json:"emailAddress,omitempty"`
	EmailAddressInvoice          string                      `json:"emailAddressInvoice,omitempty"`
	EmailAddressNewsletter       string                      `json:"emailAddressNewsletter,omitempty"`
	EmailAddressTechnicalContact string                      `json:"emailAddressTechnicalContact,omitempty"`
	Country                      string                      `json:"country,omitempty"`
	Language                     string                      `json:"language,omitempty"`
	AddDate                      string                      `json:"addDate,omitempty"`
	LastChangeDate               string                      `json:"lastChangeDate,omitempty"`
	BillingSettings              BillingSettings             `json:"billingSettings,omitempty"`
	Comments                     string                      `json:"comments,omitempty"`
	Enabled                      bool                        `json:"enabled,omitempty"`
	DnsSettings                  DnsSettings                 `json:"dnsSettings,omitempty"`
	DomainSettings               DomainSettings              `json:"domainSettings,omitempty"`
	MessageSettings              MessageSettings             `json:"messageSettings,omitempty"`
	ResellerBillingSettings      ResellerBillingSettings     `json:"resellerBillingSettings,omitempty"`
	ResellerDefaults             *ResellerDefaults           `json:"resellerDefaults"`
	ResellerWhiteLabelSettings   *ResellerWhiteLabelSettings `json:"resellerWhiteLabelSettings"`
	Rights                       []string                    `json:"rights,omitempty"`
	SalesTaxID                   string                      `json:"salesTaxID,omitempty"`
	AllowedAdvertisingMethods    []string                    `json:"allowedAdvertisingMethods,omitempty"`
	ReferralCode                 string                      `json:"referralCode,omitempty"`
	ReferringAccountId           string                      `json:"referringAccountId,omitempty"`
	Verifications                []string                    `json:"verifications,omitempty"`
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
