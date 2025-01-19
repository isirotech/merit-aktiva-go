package merit_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/egerong/merit-aktiva-go"
)

func TestGetSalesoffers(t *testing.T) {
	query := merit.GetSalesOffersQuery{
		PeriodStart: time.Now().AddDate(0, -1, 0),
		PeriodEnd:   time.Now(),
	}
	salesOffers, err := testClient.GetSalesOffers(query)
	if err != nil {
		t.Error(err)
	}
	j, err := json.MarshalIndent(salesOffers, "", "  ")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(j))
}
