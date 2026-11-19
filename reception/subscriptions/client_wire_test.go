package subscriptions

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

func TestAssignWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionSubscriptions(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.AssignResellerTierRequest
	if err := json.Unmarshal([]byte(`{
  "workspace_id": "value",
  "tier_id": "value",
  "billing_period": "monthly_period",
  "idempotency_key": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Assign(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionSubscriptions(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionSubscriptions(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/reseller/workspaces/value/subscription" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/workspaces/value/subscription")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionSubscriptions(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestAssignWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionSubscriptions(t, 422, `{
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

	var request reception.AssignResellerTierRequest
	if err := json.Unmarshal([]byte(`{
  "workspace_id": "value",
  "tier_id": "value",
  "billing_period": "monthly_period",
  "idempotency_key": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Assign(context.Background(), string("value"), &request)
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

func TestChangeWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionSubscriptions(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ChangeResellerTierRequest
	if err := json.Unmarshal([]byte(`{
  "workspace_id": "value",
  "tier_id": "value",
  "billing_period": "monthly_period",
  "idempotency_key": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Change(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionSubscriptions(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionSubscriptions(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/reception/reseller/workspaces/value/subscription" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/workspaces/value/subscription")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionSubscriptions(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestChangeWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionSubscriptions(t, 422, `{
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

	var request reception.ChangeResellerTierRequest
	if err := json.Unmarshal([]byte(`{
  "workspace_id": "value",
  "tier_id": "value",
  "billing_period": "monthly_period",
  "idempotency_key": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Change(context.Background(), string("value"), &request)
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

func TestUnassignWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionSubscriptions(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Unassign(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionSubscriptions(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionSubscriptions(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/reception/reseller/workspaces/value/subscription" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/workspaces/value/subscription")
	}
}

func TestUnassignWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionSubscriptions(t, 422, `{
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

	_, err := client.Unassign(context.Background(), string("value"))
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

func TestRestoreWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionSubscriptions(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Restore(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionSubscriptions(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionSubscriptions(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/reseller/workspaces/value/subscription/restore" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/workspaces/value/subscription/restore")
	}
}

func TestRestoreWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionSubscriptions(t, 422, `{
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

	_, err := client.Restore(context.Background(), string("value"))
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

func TestCancelPendingDowngradeWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionSubscriptions(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.CancelPendingDowngrade(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionSubscriptions(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionSubscriptions(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/reseller/workspaces/value/subscription/cancel-downgrade" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/workspaces/value/subscription/cancel-downgrade")
	}
}

func TestCancelPendingDowngradeWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionSubscriptions(t, 422, `{
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

	_, err := client.CancelPendingDowngrade(context.Background(), string("value"))
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

type wireRequestReceptionSubscriptions struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionSubscriptions records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionSubscriptions(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionSubscriptions) {
	t.Helper()
	recorded := make(chan wireRequestReceptionSubscriptions, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionSubscriptions{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionSubscriptions(t *testing.T, recorded <-chan wireRequestReceptionSubscriptions) wireRequestReceptionSubscriptions {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionSubscriptions{}
	}
}

// wireJSONContainsReceptionSubscriptions reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionSubscriptions(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionSubscriptions(got, want)
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

// setWireValueReceptionSubscriptions sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionSubscriptions(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionSubscriptions(got, want []byte) bool {
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
