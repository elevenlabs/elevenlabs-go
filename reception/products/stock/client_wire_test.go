package stock

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
	products "github.com/elevenlabs/elevenlabs-go/reception/products"
)

func TestSetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionProductsStock(t, 200, `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request products.ReceptionUpdateProductStockRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "product_id": "value",
  "location_id": "value",
  "on_hand": 123,
  "expected_on_hand": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Set(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionProductsStock(gotResponse, []byte(`{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionProductsStock(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/reception/products/value/stock" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/products/value/stock")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionProductsStock(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestSetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionProductsStock(t, 422, `{
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

	var request products.ReceptionUpdateProductStockRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "product_id": "value",
  "location_id": "value",
  "on_hand": 123,
  "expected_on_hand": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	_, err := client.Set(context.Background(), string("value"), &request)
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

func TestAdjustWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionProductsStock(t, 200, `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request products.ReceptionAdjustProductStockRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "product_id": "value",
  "location_id": "value",
  "delta": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Adjust(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionProductsStock(gotResponse, []byte(`{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`)
	}

	req := receiveWireReceptionProductsStock(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/products/value/stock/adjust" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/products/value/stock/adjust")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionProductsStock(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestAdjustWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionProductsStock(t, 422, `{
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

	var request products.ReceptionAdjustProductStockRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "product_id": "value",
  "location_id": "value",
  "delta": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	_, err := client.Adjust(context.Background(), string("value"), &request)
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

type wireRequestReceptionProductsStock struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionProductsStock records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionProductsStock(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionProductsStock) {
	t.Helper()
	recorded := make(chan wireRequestReceptionProductsStock, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionProductsStock{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionProductsStock(t *testing.T, recorded <-chan wireRequestReceptionProductsStock) wireRequestReceptionProductsStock {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionProductsStock{}
	}
}

// wireJSONContainsReceptionProductsStock reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionProductsStock(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionProductsStock(got, want)
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

// setWireValueReceptionProductsStock sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionProductsStock(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionProductsStock(got, want []byte) bool {
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
