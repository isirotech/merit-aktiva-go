package merit_test

import (
	"os"

	"github.com/egerong/merit-aktiva-go"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func NewTestClient() *merit.Client {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Failed to load .env file", zap.Error(err))
		return nil
	}
	apiID := os.Getenv("MERIT_API_ID")
	apiKey := os.Getenv("MERIT_API_KEY")
	client := merit.NewClient(apiID, apiKey, merit.API_HOST_EST, logger)
	return client
}

var testClient = NewTestClient()
