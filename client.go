package merit

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-json-experiment/json"
	"go.uber.org/zap"
)

type API_HOST string

const (
	API_HOST_EST API_HOST = "aktiva.merit.ee"
	API_HOST_FIN API_HOST = "aktiva.meritaktiva.fi"
	apiPath      string   = "api/"
)

type Client struct {
	apiID   string
	apiKey  string
	apiHost API_HOST
	logger  *zap.Logger
}

func NewClient(apiID, apiKey string, apiHost API_HOST, logger *zap.Logger) *Client {
	if logger == nil {
		logger = zap.NewNop()
	}
	return &Client{
		apiID:   apiID,
		apiKey:  apiKey,
		apiHost: apiHost,
		logger:  logger,
	}
}

// Generate a signature for the payload
func (c *Client) signature(timestamp string, payload []byte) string {
	mac := hmac.New(sha256.New, []byte(c.apiKey))
	mac.Write([]byte(c.apiID))
	mac.Write([]byte(timestamp))
	mac.Write(payload)

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// Send a HTTPS POST request with payload to the API
// Include APIID and signature in the request
func (c *Client) post(endpoint apiEndpoint, payload interface{}, dest interface{}) error {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	c.logger.Sugar().Debugf("Request body: %s", string(jsonPayload))
	timestamp := time.Now().UTC().Format("20060102150405")
	signature := c.signature(timestamp, jsonPayload)

	v := url.Values{}
	v.Set("ApiId", c.apiID)
	v.Set("timestamp", timestamp)
	v.Set("signature", signature)

	perform := url.URL{
		Scheme:   "https",
		Host:     string(c.apiHost),
		Path:     apiPath + string(endpoint),
		RawQuery: v.Encode(),
	}

	c.logger.Sugar().Debugf("Performing request to %s", perform.String())

	req, err := http.NewRequest("POST", perform.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	c.logger.Sugar().Debugf("Response body: %s", body.String())

	if resp.StatusCode != 200 {
		return fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	if dest != nil {
		err = json.Unmarshal(body.Bytes(), dest)
		if err != nil {
			return err
		}
	}

	return nil

}
