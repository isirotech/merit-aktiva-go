package merit_test

import (
	"encoding/json"
	"testing"
)

func TestGetUnitsOfMeasure(t *testing.T) {
	units, err := testClient.GetUnitsOfMeasure()
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(units, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
	if len(units) == 0 {
		t.Error("No units found")
		return
	}
}
