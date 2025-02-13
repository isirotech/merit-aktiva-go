package merit_test

import (
	"encoding/json"
	"testing"

	"github.com/egerong/merit-aktiva-go"
)

func TestGetItems(t *testing.T) {

	items, err := testClient.GetItems(merit.GetItemsQuery{})
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}

func TestGetItemsByCode(t *testing.T) {
	items, err := testClient.GetItems(merit.GetItemsQuery{
		Code: "22% Kaup",
	})
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}
