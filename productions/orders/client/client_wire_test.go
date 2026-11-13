package client

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
	"time"

	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
	productions "github.com/elevenlabs/elevenlabs-go/productions"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireProductionsOrdersClient(t, 200, `{
  "orders": [
    {
      "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
      "name": "value",
      "state": "open",
      "total_amount_usd": 9.99,
      "sandbox": true,
      "submitted_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "cancel_reason": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	offsetValue := int(123)
	startDateValue := time.Time(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	endDateValue := time.Time(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	request := &productions.ListOrdersRequest{
		PageSize:  &pageSizeValue,
		Offset:    &offsetValue,
		StartDate: &startDateValue,
		EndDate:   &endDateValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualProductionsOrdersClient(gotResponse, []byte(`{
  "orders": [
    {
      "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
      "name": "value",
      "state": "open",
      "total_amount_usd": 9.99,
      "sandbox": true,
      "submitted_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "cancel_reason": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "orders": [
    {
      "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
      "name": "value",
      "state": "open",
      "total_amount_usd": 9.99,
      "sandbox": true,
      "submitted_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "cancel_reason": "value"
    }
  ]
}`)
	}

	req := receiveWireProductionsOrdersClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/productions/orders" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/productions/orders")
	}
	if req.Query.Get("start_date") == "" {
		t.Fatalf("query %s was not sent", "start_date")
	}
	if req.Query.Get("end_date") == "" {
		t.Fatalf("query %s was not sent", "end_date")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("offset"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "offset", got, "123")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireProductionsOrdersClient(t, 422, `{
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

	pageSizeValue := int(100)
	offsetValue := int(123)
	startDateValue := time.Time(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	endDateValue := time.Time(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))
	request := &productions.ListOrdersRequest{
		PageSize:  &pageSizeValue,
		Offset:    &offsetValue,
		StartDate: &startDateValue,
		EndDate:   &endDateValue,
	}
	_, err := client.List(context.Background(), request)
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireProductionsOrdersClient(t, 200, `{
  "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "sandbox": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request productions.CreateOrderRequest
	if err := json.Unmarshal([]byte(`{
  "sandbox": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualProductionsOrdersClient(gotResponse, []byte(`{
  "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "sandbox": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "sandbox": true
}`)
	}

	req := receiveWireProductionsOrdersClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/productions/orders" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/productions/orders")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsProductionsOrdersClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireProductionsOrdersClient(t, 422, `{
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

	var request productions.CreateOrderRequest
	if err := json.Unmarshal([]byte(`{
  "sandbox": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Create(context.Background(), &request)
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

type wireRequestProductionsOrdersClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireProductionsOrdersClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireProductionsOrdersClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestProductionsOrdersClient) {
	t.Helper()
	recorded := make(chan wireRequestProductionsOrdersClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestProductionsOrdersClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireProductionsOrdersClient(t *testing.T, recorded <-chan wireRequestProductionsOrdersClient) wireRequestProductionsOrdersClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestProductionsOrdersClient{}
	}
}

// wireJSONContainsProductionsOrdersClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsProductionsOrdersClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualProductionsOrdersClient(got, want)
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

// setWireValueProductionsOrdersClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueProductionsOrdersClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualProductionsOrdersClient(got, want []byte) bool {
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
