package merit

import (
	"encoding/json"
	"testing"
	"time"
)

func TestGetSalesoffers(t *testing.T) {
	query := GetSalesOffersQuery{
		PeriodStart: time.Now().AddDate(0, -1, 0),
		PeriodEnd:   time.Now(),
	}
	salesOffers, err := testClient.GetSalesOffers(query)
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(salesOffers, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}
