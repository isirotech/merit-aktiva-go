package merit_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/egerong/merit-aktiva-go"
)

func TestInventoryMovements(t *testing.T) {
	inventoryMovements, err := testClient.GetInventoryMovements(merit.GetInventoryMovementsQuery{})
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(inventoryMovements, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}

func TestInventoryMovementsLastWeek(t *testing.T) {
	inventoryMovements, err := testClient.GetInventoryMovements(merit.GetInventoryMovementsQuery{
		PeriodStart: time.Now().AddDate(0, 0, -21),
		PeriodEnd:   time.Now(),
	})
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(inventoryMovements, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}
