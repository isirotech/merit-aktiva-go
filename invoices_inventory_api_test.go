package merit

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"go.uber.org/zap"
)

func newTestClientWithServer(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()

	ts := httptest.NewTLSServer(handler)
	t.Cleanup(ts.Close)

	host := strings.TrimPrefix(ts.URL, "https://")
	client := NewClient("test-api-id", "test-api-key", API_HOST(host), zap.NewNop())
	client.httpClient = ts.Client()

	return client, ts
}

func TestGetInvoicesByNumber_RequestAndMapping(t *testing.T) {
	client, _ := newTestClientWithServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/getinvoices2" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}

		payload := string(body)
		if !strings.Contains(payload, `"InvNo":"ARV-0001"`) {
			t.Fatalf("request payload missing InvNo, got: %s", payload)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"SIHId":"11111111-1111-1111-1111-111111111111","InvoiceNo":"ARV-0001","DocumentDate":"20260531","ProjectCode":"PRJ-1","EInvSent":true}]`))
	})

	invoices, err := client.GetInvoicesByNumber(GetInvoicesByNumberQuery{InvNo: "ARV-0001"})
	if err != nil {
		t.Fatalf("GetInvoicesByNumber returned error: %v", err)
	}

	if len(invoices) != 1 {
		t.Fatalf("expected 1 invoice, got %d", len(invoices))
	}
	if invoices[0].InvoiceNo != "ARV-0001" {
		t.Fatalf("unexpected invoice number: %s", invoices[0].InvoiceNo)
	}
	if !invoices[0].EInvSent {
		t.Fatal("expected EInvSent=true")
	}
}

func TestGetInvoiceDetails_RequestAndMapping(t *testing.T) {
	client, _ := newTestClientWithServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/getinvoice" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}

		payload := string(body)
		if !strings.Contains(payload, `"Id":"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"`) {
			t.Fatalf("request payload missing Id, got: %s", payload)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Header":{"SIHId":"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa","InvoiceNo":"ARV-0002","DocumentDate":"20260531","ProjectCode":"PRJ-2","OfferId":"bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb","OfferNo":"P-0002","EInvSent":false}}`))
	})

	details, err := client.GetInvoiceDetails(GetInvoiceDetailsQuery{
		ID:            "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
		AddAttachment: false,
	})
	if err != nil {
		t.Fatalf("GetInvoiceDetails returned error: %v", err)
	}

	if details.Header.InvoiceNo != "ARV-0002" {
		t.Fatalf("unexpected invoice number: %s", details.Header.InvoiceNo)
	}
	if details.Header.OfferNo != "P-0002" {
		t.Fatalf("unexpected offer number: %s", details.Header.OfferNo)
	}
	if details.Header.EInvSent {
		t.Fatal("expected EInvSent=false")
	}
}

func TestSendInventoryMovement_SerializesTypeOutAndDate(t *testing.T) {
	client, _ := newTestClientWithServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/SendInvMovement" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}

		payload := string(body)
		if !strings.Contains(payload, `"Type":2`) {
			t.Fatalf("expected writeoff type 2 in payload, got: %s", payload)
		}
		if !strings.Contains(payload, `"DocDate":"20260531"`) {
			t.Fatalf("expected formatted DocDate in payload, got: %s", payload)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	})

	err := client.SendInventoryMovement(SendInventoryMovementQuery{
		DocDate:       time.Date(2026, 5, 31, 0, 0, 0, 0, time.UTC),
		DocNo:         "WR-001",
		Location1Code: "MAIN",
		Type:          InventoryMovementTypeOut,
		Rows: []SendInventoryRow{
			{
				ArticleCode: "ART-001",
				UOMName:     "pcs",
				Quantity:    3,
			},
		},
	})
	if err != nil {
		t.Fatalf("SendInventoryMovement returned error: %v", err)
	}
}

func TestSendInventoryMovement_Non200ReturnsError(t *testing.T) {
	client, _ := newTestClientWithServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid movement"))
	})

	err := client.SendInventoryMovement(SendInventoryMovementQuery{
		DocDate:       time.Date(2026, 5, 31, 0, 0, 0, 0, time.UTC),
		Location1Code: "MAIN",
		Type:          InventoryMovementTypeOut,
		Rows: []SendInventoryRow{{
			ArticleCode: "ART-001",
			Quantity:    1,
		}},
	})
	if err == nil {
		t.Fatal("expected error for non-200 response, got nil")
	}
	if !strings.Contains(err.Error(), "status code 400") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUpdateItem_RequestAndEndpoint(t *testing.T) {
	client, _ := newTestClientWithServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/updateitem" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}

		payload := string(body)
		if !strings.Contains(payload, `"Id":"aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"`) {
			t.Fatalf("request payload missing Id, got: %s", payload)
		}
		if !strings.Contains(payload, `"NonActive":true`) {
			t.Fatalf("request payload missing NonActive=true, got: %s", payload)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	})

	itemID, err := guid.FromString("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	if err != nil {
		t.Fatalf("invalid guid for test: %v", err)
	}

	err = client.UpdateItem(UpdateItemQuery{
		ID:          itemID,
		Code:        "P-100558",
		Description: "Andrii DC",
		NonActive:   true,
	})
	if err != nil {
		t.Fatalf("UpdateItem returned error: %v", err)
	}
}
