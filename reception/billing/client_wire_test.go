package billing

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
	reception "github.com/elevenlabs/elevenlabs-go/reception"
)

func TestListInvoicesWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionBilling(t, 200, `{
  "invoices": [
    {
      "invoice_id": "value",
      "workspace_id": "value",
      "workspace_name": "value",
      "status": "paid",
      "amount_cents": 123,
      "currency": "value",
      "created_unix": 123,
      "paid_unix": 123,
      "pay_url": "value",
      "pdf_url": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(100)
	request := &reception.ListInvoicesBillingRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
	}
	result, err := client.ListInvoices(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionBilling(gotResponse, []byte(`{
  "invoices": [
    {
      "invoice_id": "value",
      "workspace_id": "value",
      "workspace_name": "value",
      "status": "paid",
      "amount_cents": 123,
      "currency": "value",
      "created_unix": 123,
      "paid_unix": 123,
      "pay_url": "value",
      "pdf_url": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "invoices": [
    {
      "invoice_id": "value",
      "workspace_id": "value",
      "workspace_name": "value",
      "status": "paid",
      "amount_cents": 123,
      "currency": "value",
      "created_unix": 123,
      "paid_unix": 123,
      "pay_url": "value",
      "pdf_url": "value"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireReceptionBilling(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/reseller/billing/invoices" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/billing/invoices")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
}

func TestListInvoicesWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionBilling(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	cursorValue := string("value")
	pageSizeValue := int(100)
	request := &reception.ListInvoicesBillingRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
	}
	_, err := client.ListInvoices(context.Background(), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestGetInvoiceWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionBilling(t, 200, `{
  "invoice_id": "value",
  "workspace_id": "value",
  "workspace_name": "value",
  "status": "paid",
  "amount_cents": 123,
  "currency": "value",
  "created_unix": 123,
  "paid_unix": 123,
  "pay_url": "value",
  "pdf_url": "value",
  "lines": [
    {
      "kind": "plan",
      "description": "value",
      "quantity": 123,
      "amount_cents": 123,
      "period_start_unix": 123,
      "period_end_unix": 123
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetInvoice(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionBilling(gotResponse, []byte(`{
  "invoice_id": "value",
  "workspace_id": "value",
  "workspace_name": "value",
  "status": "paid",
  "amount_cents": 123,
  "currency": "value",
  "created_unix": 123,
  "paid_unix": 123,
  "pay_url": "value",
  "pdf_url": "value",
  "lines": [
    {
      "kind": "plan",
      "description": "value",
      "quantity": 123,
      "amount_cents": 123,
      "period_start_unix": 123,
      "period_end_unix": 123
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "invoice_id": "value",
  "workspace_id": "value",
  "workspace_name": "value",
  "status": "paid",
  "amount_cents": 123,
  "currency": "value",
  "created_unix": 123,
  "paid_unix": 123,
  "pay_url": "value",
  "pdf_url": "value",
  "lines": [
    {
      "kind": "plan",
      "description": "value",
      "quantity": 123,
      "amount_cents": 123,
      "period_start_unix": 123,
      "period_end_unix": 123
    }
  ]
}`)
	}

	req := receiveWireReceptionBilling(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/reseller/billing/invoices/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/billing/invoices/value")
	}
}

func TestGetInvoiceWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionBilling(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.GetInvoice(context.Background(), string("value"))
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestListUpcomingPaymentsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionBilling(t, 200, `{
  "payments": [
    {
      "workspace_id": "value",
      "workspace_name": "value",
      "subscription": {
        "tier_id": "value",
        "billing_period": "monthly_period",
        "billing_period_end_unix": 123,
        "cancel_at_unix": 123,
        "pending_downgrade_tier": "value",
        "pending_downgrade_at_unix": 123
      },
      "plan_usd": 9.99,
      "overage_so_far_usd": 9.99
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.ListUpcomingPayments(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionBilling(gotResponse, []byte(`{
  "payments": [
    {
      "workspace_id": "value",
      "workspace_name": "value",
      "subscription": {
        "tier_id": "value",
        "billing_period": "monthly_period",
        "billing_period_end_unix": 123,
        "cancel_at_unix": 123,
        "pending_downgrade_tier": "value",
        "pending_downgrade_at_unix": 123
      },
      "plan_usd": 9.99,
      "overage_so_far_usd": 9.99
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "payments": [
    {
      "workspace_id": "value",
      "workspace_name": "value",
      "subscription": {
        "tier_id": "value",
        "billing_period": "monthly_period",
        "billing_period_end_unix": 123,
        "cancel_at_unix": 123,
        "pending_downgrade_tier": "value",
        "pending_downgrade_at_unix": 123
      },
      "plan_usd": 9.99,
      "overage_so_far_usd": 9.99
    }
  ]
}`)
	}

	req := receiveWireReceptionBilling(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/reseller/billing/upcoming" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/billing/upcoming")
	}
}

func TestListUpcomingPaymentsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionBilling(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.ListUpcomingPayments(context.Background())
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

type wireRequestReceptionBilling struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionBilling records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionBilling(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionBilling) {
	t.Helper()
	recorded := make(chan wireRequestReceptionBilling, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionBilling{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
		default:
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if responseBody != "" {
			_, _ = w.Write([]byte(responseBody))
		}
	}))
	t.Cleanup(server.Close)
	return server, recorded
}

func receiveWireReceptionBilling(t *testing.T, recorded <-chan wireRequestReceptionBilling) wireRequestReceptionBilling {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionBilling{}
	}
}

// wireJSONContainsReceptionBilling reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionBilling(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionBilling(got, want)
	}
	for key, wantValue := range wantObject {
		if key == "stream" {
			continue
		}
		if gotValue, ok := gotObject[key]; !ok || !reflect.DeepEqual(gotValue, wantValue) {
			return false
		}
	}
	return true
}

// setWireValueReceptionBilling sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionBilling(t *testing.T, target any, value any) {
	t.Helper()
	field := reflect.ValueOf(target).Elem()
	if field.Kind() == reflect.Ptr {
		field.Set(reflect.New(field.Type().Elem()))
		field = field.Elem()
	}
	switch typed := value.(type) {
	case string:
		field.SetString(typed)
	case int:
		field.SetInt(int64(typed))
	case float64:
		field.SetFloat(typed)
	default:
		t.Fatalf("unsupported value %T", value)
	}
}

func wireJSONBytesEqualReceptionBilling(got, want []byte) bool {
	var gotValue any
	var wantValue any
	if err := json.Unmarshal(got, &gotValue); err != nil {
		return false
	}
	if err := json.Unmarshal(want, &wantValue); err != nil {
		return false
	}
	return reflect.DeepEqual(gotValue, wantValue)
}
