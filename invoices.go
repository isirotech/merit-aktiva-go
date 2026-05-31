package merit

import (
	"fmt"
	"strings"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
)

// GetInvoicesByNumberQuery is the request payload for v2/getinvoices2.
type GetInvoicesByNumberQuery struct {
	InvNo    string    `json:"InvNo,omitempty"`
	CustName string    `json:"CustName,omitempty"`
	CustID   guid.GUID `json:"CustId,omitzero"`
}

// SalesInvoiceSummary contains top-level invoice fields returned by v2/getinvoices2.
type SalesInvoiceSummary struct {
	SIHID        guid.GUID `json:"SIHId"`
	InvoiceNo    string    `json:"InvoiceNo"`
	DocumentDate string    `json:"DocumentDate"`
	ProjectCode  string    `json:"ProjectCode"`
	EInvSent     bool      `json:"EInvSent"`
}

// GetInvoicesQuery is the request payload for v2/getinvoices.
type GetInvoicesQuery struct {
	PeriodStart time.Time `json:"PeriodStart,omitempty"`
	PeriodEnd   time.Time `json:"PeriodEnd,omitempty"`
	UnPaid      bool      `json:"UnPaid,omitempty"`
	DateType    int       `json:"DateType,omitempty"`
}

type getInvoicesQueryFormated struct {
	PeriodStart queryDate `json:"PeriodStart,omitempty"`
	PeriodEnd   queryDate `json:"PeriodEnd,omitempty"`
	UnPaid      bool      `json:"UnPaid,omitempty"`
	DateType    int       `json:"DateType,omitempty"`
}

func (q GetInvoicesQuery) format() getInvoicesQueryFormated {
	return getInvoicesQueryFormated{
		PeriodStart: queryDate{q.PeriodStart, "20060102"},
		PeriodEnd:   queryDate{q.PeriodEnd, "20060102"},
		UnPaid:      q.UnPaid,
		DateType:    q.DateType,
	}
}

// GetInvoices returns invoice summaries for v2/getinvoices period lookups.
func (c *Client) GetInvoices(query GetInvoicesQuery) ([]SalesInvoiceSummary, error) {
	queryFormated := query.format()
	invoices := []SalesInvoiceSummary{}
	err := c.post(epGetSalesInvoicesByPeriod, queryFormated, &invoices)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

// GetInvoicesByNumber returns invoice summaries for v2/getinvoices2 lookups.
func (c *Client) GetInvoicesByNumber(query GetInvoicesByNumberQuery) ([]SalesInvoiceSummary, error) {
	invoices := []SalesInvoiceSummary{}
	err := c.post(epGetSalesInvoicesByNumber, query, &invoices)
	if err != nil {
		return nil, err
	}

	return invoices, nil
}

// GetInvoiceDetailsQuery is the request payload for v2/getinvoice.
type GetInvoiceDetailsQuery struct {
	ID            string `json:"Id"`
	AddAttachment bool   `json:"AddAttachment,omitempty"`
}

// InvoiceHeader contains header fields returned by v2/getinvoice.
type InvoiceHeader struct {
	SIHID        guid.GUID `json:"SIHId"`
	InvoiceNo    string    `json:"InvoiceNo"`
	DocumentDate string    `json:"DocumentDate"`
	ProjectCode  string    `json:"ProjectCode"`
	OfferID      guid.GUID `json:"OfferId"`
	OfferNo      string    `json:"OfferNo"`
	EInvSent     bool      `json:"EInvSent"`
}

// InvoiceDimensionAllocation contains line-level dimension allocation values.
type InvoiceDimensionAllocation struct {
	Code string `json:"Code"`
}

// InvoiceRow contains invoice line fields needed by merit-server workflows.
type InvoiceRow struct {
	DimAllocation     []InvoiceDimensionAllocation `json:"DimAllocation"`
	ProjectAllocation []InvoiceDimensionAllocation `json:"ProjectAllocation"`
}

// InvoiceDetails contains the response payload for v2/getinvoice.
type InvoiceDetails struct {
	Header InvoiceHeader `json:"Header"`
	Lines  []InvoiceRow  `json:"Lines"`
}

// GetInvoiceDetails returns invoice details for a specific sales invoice id.
func (c *Client) GetInvoiceDetails(query GetInvoiceDetailsQuery) (*InvoiceDetails, error) {
	var details InvoiceDetails
	err := c.post(epGetSalesInvoiceDetails, query, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

// SendInvoiceAsEInvQuery is the request payload for v2/sendinvoiceaseinv.
type SendInvoiceAsEInvQuery struct {
	ID        string `json:"Id"`        // SIHId GUID of the sales invoice
	DelivNote bool   `json:"DelivNote"` // true = delivery note format (omits prices)
}

// SendInvoiceAsEInv sends a sales invoice as an e-invoice via Merit Aktiva.
// Returns (true, nil) on success, including when Merit sends without operator routing ("api-noeinv").
// Returns (true, nil) when the invoice was already sent as an e-invoice (idempotent, HTTP 400).
// Returns (false, err) on transport or unexpected response errors.
func (c *Client) SendInvoiceAsEInv(query SendInvoiceAsEInvQuery) (bool, error) {
	status, raw, err := c.postRaw(epSendInvoiceAsEInv, query)
	if err != nil {
		return false, err
	}
	switch status {
	case 200:
		switch strings.TrimSpace(string(raw)) {
		case "OK", "api-noeinv":
			// "api-noeinv" means Merit sent the e-invoice without operator routing —
			// the invoice is in the desired sent state.
			return true, nil
		default:
			return false, fmt.Errorf("unexpected response from sendinvoiceaseinv: %q", strings.TrimSpace(string(raw)))
		}
	case 400:
		// Merit returns 400 when the invoice has already been sent as an e-invoice.
		// Treat as idempotent success — the invoice is in the desired sent state.
		return true, nil
	default:
		return false, fmt.Errorf("API returned status %d: %s", status, strings.TrimSpace(string(raw)))
	}
}
