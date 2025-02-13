package merit_test

import (
	"encoding/json"
	"testing"

	"github.com/egerong/merit-aktiva-go"
)

func TestGetTaxes(t *testing.T) {
	taxes, err := testClient.GetTaxes(nil)
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(taxes, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}

func TestGetTaxByCode(t *testing.T) {
	query := merit.Tax{
		Code: "22%",
	}
	tax, err := testClient.GetTaxes(&query)
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(tax, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
	if len(tax) == 0 {
		t.Error("No tax found")
		return
	}
	if len(tax) > 1 {
		t.Error("Multiple taxes found")
		return
	}
}
