package merit

import (
	"fmt"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	"github.com/shopspring/decimal"
)

const (
	OfferTypeQuote OfferType = iota + 1
	OfferTypeSalesOrder
	OfferTypePrepaymentInvoice
)

type OfferStatus int

const (
	OfferStatusCreated OfferStatus = iota + 1
	OfferStatusSent
	OfferStatusApproved
	OfferStatusRejected
	OfferStatusCommentReceived
	OfferStatusInvoiceCreated
	OfferStatusCanceled
)

func (s OfferStatus) String() string {
	switch s {
	case OfferStatusCreated:
		return "Created"
	case OfferStatusSent:
		return "Sent"
	case OfferStatusApproved:
		return "Approved"
	case OfferStatusRejected:
		return "Rejected"
	case OfferStatusCommentReceived:
		return "Comment Received"
	case OfferStatusInvoiceCreated:
		return "Invoice Created"
	case OfferStatusCanceled:
		return "Canceled"
	default:
		return "Unknown"
	}
}

type SalesOffer struct {
	ID              string          `json:"SIHId"`
	DepartmentName  string          `json:"DepartmentName"`
	Dimension1Code  string          `json:"Dimension1Code"`
	Dimension2Code  string          `json:"Dimension2Code"`
	Dimension3Code  string          `json:"Dimension3Code"`
	Dimension4Code  string          `json:"Dimension4Code"`
	Dimension5Code  string          `json:"Dimension5Code"`
	Dimension6Code  string          `json:"Dimension6Code"`
	Dimension7Code  string          `json:"Dimension7Code"`
	BatchInfo       string          `json:"BatchInfo"`
	OfferNo         string          `json:"OfferNo"`
	DocumentDate    string          `json:"DocumentDate"`
	TransactionDate string          `json:"TransactionDate"`
	CustomerId      string          `json:"CustomerId"`
	CustomerName    string          `json:"CustomerName"`
	HComment        string          `json:"HComment"`
	FComment        string          `json:"FComment"`
	DueDate         string          `json:"DueDate"`
	CurrencyCode    string          `json:"CurrencyCode"`
	CurrencyRate    decimal.Decimal `json:"CurrencyRate"`
	TaxAmount       decimal.Decimal `json:"TaxAmount"`
	RoundingAmount  decimal.Decimal `json:"RoundingAmount"`
	TotalAmount     decimal.Decimal `json:"TotalAmount"`
	ProfitAmount    decimal.Decimal `json:"ProfitAmount"`
	TotalSum        decimal.Decimal `json:"TotalSum"`
	UserName        string          `json:"UserName"`
	ReferenceNo     string          `json:"ReferenceNo"`
	PriceInclVat    bool            `json:"PriceInclVat"`
	VatRegNo        string          `json:"VatRegNo"`
	PaidAmount      decimal.Decimal `json:"PaidAmount"`
	EInvSent        bool            `json:"EInvSent"`
	EmailSent       string          `json:"EmailSent"`
	DocType         OfferType       `json:"DocType"`
	DocStatus       OfferStatus     `json:"DocStatus"`
	DeliveryDate    string          `json:"DeliveryDate"`
	Paid            bool            `json:"Paid"`
	ChangedDate     string          `json:"ChangedDate"`
}

type GetSalesOffersQuery struct {
	PeriodStart time.Time `json:"PeriodStart,omitempty"`
	PeriodEnd   time.Time `json:"PeriodEnd,omitempty"`
	DateType    int       `json:"DateType,omitempty"`
	UnPaid      bool      `json:"UnPaid,omitempty"`
}

type getSalesOffersQueryFormated struct {
	PeriodStart queryDate `json:"PeriodStart,omitempty"`
	PeriodEnd   queryDate `json:"PeriodEnd,omitempty"`
	DateType    int       `json:"DateType,omitempty"`
	UnPaid      bool      `json:"UnPaid,omitempty"`
}

func (query GetSalesOffersQuery) format() getSalesOffersQueryFormated {
	return getSalesOffersQueryFormated{
		PeriodStart: queryDate{query.PeriodStart, "20060102"},
		PeriodEnd:   queryDate{query.PeriodEnd, "20060102"},
		DateType:    query.DateType,
		UnPaid:      query.UnPaid,
	}
}

func (c *Client) GetSalesOffers(query GetSalesOffersQuery) ([]SalesOffer, error) {
	queryFormated := query.format()
	salesOffers := []SalesOffer{}
	err := c.post(epGetListOfSalesOffers, queryFormated, &salesOffers)
	if err != nil {
		return nil, err
	}
	return salesOffers, nil
}

type OfferType int

type CreateSalesOfferQuery struct {
	Customer       CustomerObject // CustomerObject
	DocDate        time.Time
	ExpireDate     time.Time // if DocType 2 or 3, ExpireDate=DueDate
	DeliveryDate   time.Time
	OfferNo        string      `json:"OfferNo"`            // Required
	DocType        OfferType   `json:"DocType,omitzero"`   // 1=quote, 2=sales order, 3=prepayment invoice
	DocStatus      OfferStatus `json:"DocStatus,omitzero"` // 1=created, 2=sent, 3=approved, 4=rejected, 5=comment received, 6=invoice created, 7=canceled
	RefNo          string      `json:"RefNo,omitempty"`    // Please validate this number yourself.
	CurrencyCode   string      `json:"CurrencyCode,omitempty"`
	DepartmentCode string      `json:"DepartmentCode,omitempty"` // If used then must be found in the company database.
	Dimensions     []DimensionsObject
	OfferRow       []OfferRow
	TotalAmount    *decimal.Decimal `json:"TotalAmount,omitempty"`    // Amount without VAT
	RoundingAmount *decimal.Decimal `json:"RoundingAmount,omitempty"` // Use it for getting PDF invoice to round number. Does not affect TotalAmount.
	TaxAmount      []TaxObject      // Required
	Payment        *OfferPayment    `json:"Payment,omitempty"`
	HeaderComment  string           `json:"Hcomment,omitempty"`     // If not specified, API will get it from client record, if it is written there.
	FooterComment  string           `json:"Fcomment,omitempty"`     // If not specified, API will get it from client record, if it is written there.
	ReserveItems   bool             `json:"ReserveItems,omitempty"` // If true, then stock items will be reserved.
	PrepaymPct     *decimal.Decimal `json:"PrepaymPct,omitempty"`   // Prepayment percentage. Required for prepayment invoices.
	Payer          *PayerObject     `json:"Payer,omitempty"`
}

type OfferRow struct {
	Item           ItemObject       `json:"Item"` // Sometimes the volume of transactions in the sales software is very high and there is no need to duplicate all the data in accounting. In those cases, you could consider using the same item code for the items with the same VAT rate.
	Quantity       decimal.Decimal  `json:"Quantity"`
	Price          decimal.Decimal  `json:"Price"`
	DiscountPct    *decimal.Decimal `json:"DiscountPct,omitempty"`
	DiscountAmount *decimal.Decimal `json:"DiscountAmount,omitempty"` // Amount * Price * (DiscountPCt / 100). This is not rounded. Will be substracted from row amount before row roundings.
	TaxId          string           `json:"TaxId,omitempty"`
	LocationCode   string           `json:"LocationCode,omitempty"`   // Used for stock items and multiple stocks. If used then must be found in the company database.
	DepartmentCode string           `json:"DepartmentCode,omitempty"` // If used then must be found in the company database.
	ItemCostAmount *decimal.Decimal `json:"ItemCostAmount,omitempty"` // Required for credit invoices when crediting stock items.
	GLAccountCode  string           `json:"GLAccountCode,omitempty"`  // If used, must be found in the company database.
	ProjectCode    string           `json:"ProjectCode,omitempty"`    // If used, must be found in the company database.
	CostCenterCode string           `json:"CostCenterCode,omitempty"` // If used, must be found in the company database.
}

type OfferPayment struct {
	PaymentMethod string          `json:"PaymentMethod"` // Name of the payment method. Must be found in the company database.
	PaidAmount    decimal.Decimal `json:"PaidAmount"`    // Amount with VAT (not more) or less if partial payment
	PaymDate      string          `json:"PaymDate"`      // YYYYmmddHHii
}

type PayerObject struct{}

// CustomerId
// InvoiceId
// InvoiceNo
// RefNo
// NewCustomer

type CreateSalesOfferResponse struct {
	CustomerID  guid.GUID `json:"CustomerId"`
	InvoiceID   guid.GUID `json:"InvoiceId"`
	InvoiceNo   string    `json:"InvoiceNo"`
	RefNo       string    `json:"RefNo"`
	NewCustomer bool      `json:"NewCustomer"`
}

type CreateSalesOfferQueryFormated struct {
	CreateSalesOfferQuery
	DocDate      queryDate `json:"DocDate,omitzero"`
	ExpireDate   queryDate `json:"ExpireDate,omitzero"`
	DeliveryDate queryDate `json:"DeliveryDate,omitzero"`
}

func (c *Client) CreateSalesOffer(query CreateSalesOfferQuery) (*CreateSalesOfferResponse, error) {
	queryFormated := CreateSalesOfferQueryFormated{
		CreateSalesOfferQuery: query,
		DocDate:               queryDate{query.DocDate, "20060102"},
		ExpireDate:            queryDate{query.ExpireDate, "20060102"},
		DeliveryDate:          queryDate{query.DeliveryDate, "20060102"},
	}
	j, err := json.Marshal(queryFormated)
	if err != nil {
		return nil, err
	}
	(*jsontext.Value)(&j).Indent("", "  ") // indent for readability
	fmt.Print(string(j))
	var response CreateSalesOfferResponse
	err = c.post(epCreateSalesOffer, queryFormated, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// SetOfferStatusQuery is the payload for the v2/setofferstatus endpoint.
type SetOfferStatusQuery struct {
	ID        string      `json:"ID"`
	NewStatus OfferStatus `json:"NewStatus"`
	Comment   string      `json:"Comment,omitempty"` // Required if NewStatus == OfferStatusCommentReceived
}

// SetOfferStatus changes the status of an existing sales offer.
func (c *Client) SetOfferStatus(query SetOfferStatusQuery) error {
	return c.post(epSetOfferStatus, query, nil)
}
