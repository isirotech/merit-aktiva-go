package merit

import (
	"encoding/json"
	"testing"
)

func TestGetCustomers(t *testing.T) {
	query := GetCustomersQuery{}
	customers, err := testClient.GetCustomers(query)
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}
