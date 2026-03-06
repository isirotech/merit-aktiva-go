package merit

import (
	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/shopspring/decimal"
)

type TaxObject struct {
	TaxID  guid.GUID       `json:"TaxId"`  // Required. Use gettaxes endpoint to detect the guid needed
	Amount decimal.Decimal `json:"Amount"` // Required
}

// Field Name	Type	Comment
// Id	Guid
// Code	Str 16
// Name	Str 40
// NameEN	Str 40
// NameRU	Str 40
// TaxPct	Decimal 2.2

type Tax struct {
	ID     guid.GUID       `json:"Id"`
	Code   string          `json:"Code"`
	Name   string          `json:"Name"`
	NameEN string          `json:"NameEN"`
	NameRU string          `json:"NameRU"`
	TaxPct decimal.Decimal `json:"TaxPct"`
}

func (c *Client) GetTaxes(query *Tax) ([]Tax, error) {
	taxes := []Tax{}
	err := c.post(epGetListOfTaxes, struct{}{}, &taxes)
	if err != nil {
		return nil, err
	}
	if query == nil {
		return taxes, nil
	}
	filteredTaxes := []Tax{}
	var emptyGUID guid.GUID
	var emptyDecimal decimal.Decimal
	for _, tax := range taxes {
		if query.ID != emptyGUID && tax.ID != query.ID {
			continue
		}
		if query.Code != "" && tax.Code != query.Code {
			continue
		}
		if query.Name != "" && tax.Name != query.Name {
			continue
		}
		if query.NameEN != "" && tax.NameEN != query.NameEN {
			continue
		}
		if query.NameRU != "" && tax.NameRU != query.NameRU {
			continue
		}
		if query.TaxPct != emptyDecimal && tax.TaxPct != query.TaxPct {
			continue
		}
		filteredTaxes = append(filteredTaxes, tax)
	}
	return filteredTaxes, nil
}
