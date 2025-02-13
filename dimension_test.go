package merit_test

import (
	"encoding/json"
	"testing"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/egerong/merit-aktiva-go"
)

func TestGetDimensions(t *testing.T) {
	dimensions, err := testClient.GetDimensions(merit.Dimension{})
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(dimensions, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}

func TestGetDimensionsByID(t *testing.T) {
	g, _ := guid.FromString("")
	dimensions, err := testClient.GetDimensions(merit.Dimension{
		ID: g,
	})
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(dimensions, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}

func TestGetDimensionsByCode(t *testing.T) {
	dimensions, err := testClient.GetDimensions(merit.Dimension{
		DimName: "Projekt",
		Code:    "RH233303",
	})
	if err != nil {
		t.Error(err)
		return
	}
	j, err := json.MarshalIndent(dimensions, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(j))
}
