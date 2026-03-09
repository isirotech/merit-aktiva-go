package merit

import (
	"fmt"
	"strings"
)

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
