package merit

import (
	// "merit"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/shopspring/decimal"
)

type GetCustomersQuery struct {
	ID           string `json:"Id,omitempty"`       // If filled, next fields will be ignored
	RegNo        string `json:"RegNo,omitempty"`    // Exact match.
	VatRegNo     string `json:"VatRegNo,omitempty"` // Exact match.
	Name         string `json:"Name,omitempty"`     // Broad match.
	WithComments bool   `json:"WithComments"`
	CommentsFrom string `json:"CommentsFrom,omitempty"`
	ChangedDate  string `json:"ChangedDate,omitempty"`
}

func (c *Client) GetCustomers(query GetCustomersQuery) ([]Customer, error) {
	customers := []Customer{}
	err := c.post(epGetListOfCustomers, query, &customers)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

type Customer struct {
	ID                guid.GUID          `json:"CustomerId"`
	Name              string             `json:"Name"`
	RegNo             string             `json:"RegNo"`
	Contact           string             `json:"Contact"`
	PhoneNo           string             `json:"PhoneNo"`
	Address           string             `json:"Address"`
	City              string             `json:"City"`
	Email             string             `json:"Email"`
	CurrencyCode      string             `json:"CurrencyCode"`
	CustomerGroupId   string             `json:"CustomerGroupId"`
	CustomerGroupName string             `json:"CustomerGroupName"`
	PostalCode        string             `json:"PostalCode"`
	CountryName       string             `json:"CountryName"`
	CountryCode       string             `json:"CountryCode"`
	County            string             `json:"County"`
	PhoneNo2          string             `json:"PhoneNo2"`
	FaxNo             string             `json:"FaxNo"`
	HomePage          string             `json:"HomePage"`
	PaymentDeadLine   int                `json:"PaymentDeadLine"`
	OverdueCharge     decimal.Decimal    `json:"OverdueCharge"`
	VatRegNo          string             `json:"VatRegNo"`
	NotTDCustomer     bool               `json:"NotTDCustomer"`
	BankName          string             `json:"BankName"`
	BankAccount       string             `json:"BankAccount"`
	SalesInvLang      string             `json:"SalesInvLang"`
	RefNoBase         string             `json:"RefNoBase"`
	Comments          []CommentsObject   `json:"Comments"`
	Dimensions        []DimensionsObject `json:"Dimensions"`
	ChangedDate       string             `json:"ChangedDate"`
}

type CommentsObject struct {
	CommDate string `json:"CommDate"`
	Comment  string `json:"Comment"`
}

type CustomerObject struct {
	ID              guid.GUID        `json:"Id,omitempty"`   // If filled and customer is found in the database then following fields are not important. If not found, the customer is added using the following fields.
	Name            string           `json:"Name,omitempty"` // Required when customer is added
	RegNo           string           `json:"RegNo,omitempty"`
	NotTDCustomer   bool             `json:"NotTDCustomer,omitempty"` // Required when customer is added. EE True for physical persons and foreign companies. PL True for physical persons. Allowed “true” or “false” (lowercase).
	VatRegNo        string           `json:"VatRegNo,omitempty"`
	CurrencyCode    string           `json:"CurrencyCode,omitempty"`
	PaymentDeadLine int              `json:"PaymentDeadLine,omitzero"` // If missing then taken from default settings.
	OverDueCharge   *decimal.Decimal `json:"OverDueCharge,omitempty"`  // If missing then taken from default settings.
	Address         string           `json:"Address,omitempty"`
	City            string           `json:"City,omitempty"`
	County          string           `json:"County,omitempty"`
	PostalCode      string           `json:"PostalCode,omitempty"`
	CountryCode     string           `json:"CountryCode,omitempty"` // Required when adding
	PhoneNo         string           `json:"PhoneNo,omitempty"`
	PhoneNo2        string           `json:"PhoneNo2,omitempty"`
	HomePage        string           `json:"HomePage,omitempty"`
	Email           string           `json:"Email,omitempty"`
	SalesInvLang    string           `json:"SalesInvLang,omitempty"` // Invoice language for this specific customer.(ET,EN,RU,FI,PL,SV)
	GLNCode         string           `json:"GLNCode,omitempty"`
	PartyCode       string           `json:"PartyCode,omitempty"`
	RefNoBase       string           `json:"RefNoBase,omitempty"`
	EInvPaymId      string           `json:"EInvPaymId,omitempty"`
	EInvOperator    int              `json:"EInvOperator,omitzero"` // 1 - Not exist, 2 - E-invoices to the bank through Omniva, 3 - Bank ( full extent E-invoice), 4- Bank (limited extent E-invoice)
	BankAccount     string           `json:"BankAccount,omitempty"`
}
