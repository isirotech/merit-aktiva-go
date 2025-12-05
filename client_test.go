package merit

import (
	"os"
	"testing"

	"github.com/go-json-experiment/json"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func NewTestClient() *Client {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
		return nil
	}
	apiID := os.Getenv("MERIT_API_ID")
	apiKey := os.Getenv("MERIT_API_KEY")
	client := NewClient(apiID, apiKey, API_HOST_EST, logger)
	return client
}

var testClient = NewTestClient()

func TestSignature(t *testing.T) {
	if testClient == nil {
		t.Fatal("Failed to create test client")
	}

	apiID := "670fe52f-558a-4be8-ade0-526e01a106d0"
	apiKey := "AoCmZGUfWMMhLJ+Eb6oRF4pAEw9XJP9b/RL5c2Gqk2w="
	timestamp := "20240624205902"
	data := struct {
		CustName    string `json:"CustName"`
		CustID      string `json:"CustId"`
		OverDueDays int    `json:"OverDueDays"`
		DebtDate    string `json:"DebtDate"`
	}{
		CustName:    "Kliendinimi",
		CustID:      "3a274294-9c60-4a3d-93f0-1874253f073e",
		OverDueDays: 5,
		DebtDate:    "20220501",
	}
	targetSignature := "dt6dkfuj+OfX01YkvvAoN/fekAUGr6AvVlQhUUja9Qc="

	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	client := NewClient(apiID, apiKey, API_HOST_EST, zap.NewNop())
	signature := client.signature(timestamp, payload)

	if signature != targetSignature {
		t.Errorf("Signature mismatch. Got %s, expected %s", signature, targetSignature)
	}
}
