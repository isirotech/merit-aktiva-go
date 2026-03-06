package merit

import (
	"encoding/json"
	"testing"
)

func TestGetItems(t *testing.T) {

	items, err := testClient.GetItems(GetItemsQuery{})
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
	items, err := testClient.GetItems(GetItemsQuery{
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
