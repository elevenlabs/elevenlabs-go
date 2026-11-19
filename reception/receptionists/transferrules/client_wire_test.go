package transferrules

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
	receptionists "github.com/elevenlabs/elevenlabs-go/reception/receptionists"
)

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionReceptionistsTransferrules(t, 200, `{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0",
  "id": "value",
  "created_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request receptionists.ReceptionCreateTransferRuleRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "agent_id": "value",
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Create(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionReceptionistsTransferrules(gotResponse, []byte(`{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0",
  "id": "value",
  "created_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0",
  "id": "value",
  "created_at_unix": 123
}`)
	}

	req := receiveWireReceptionReceptionistsTransferrules(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/customer-facing-agents/value/transfer-rules" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/customer-facing-agents/value/transfer-rules")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionReceptionistsTransferrules(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionReceptionistsTransferrules(t, 422, `{
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

	var request receptionists.ReceptionCreateTransferRuleRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "agent_id": "value",
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	_, err := client.Create(context.Background(), string("value"), &request)
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionReceptionistsTransferrules(t, 200, `{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0",
  "id": "value",
  "created_at_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request receptionists.ReceptionUpdateTransferRuleRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "rule_id": "value",
  "agent_id": "value",
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Update(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionReceptionistsTransferrules(gotResponse, []byte(`{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0",
  "id": "value",
  "created_at_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0",
  "id": "value",
  "created_at_unix": 123
}`)
	}

	req := receiveWireReceptionReceptionistsTransferrules(t, recorded)
	if req.Method != "PUT" {
		t.Fatalf("method: got %q want %q", req.Method, "PUT")
	}
	if req.Path != "/v1/reception/customer-facing-agents/value/transfer-rules/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/customer-facing-agents/value/transfer-rules/value2")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionReceptionistsTransferrules(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionReceptionistsTransferrules(t, 422, `{
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

	var request receptionists.ReceptionUpdateTransferRuleRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "rule_id": "value",
  "agent_id": "value",
  "label": "value",
  "phone_number": "value",
  "condition": "value",
  "post_dial_digits": "0"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	_, err := client.Update(context.Background(), string("value"), string("value2"), &request)
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

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionReceptionistsTransferrules(t, 200, `{
  "key": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &receptionists.DeleteTransferRulesRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.Delete(context.Background(), string("value"), string("value2"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionReceptionistsTransferrules(gotResponse, []byte(`{
  "key": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": "value"
}`)
	}

	req := receiveWireReceptionReceptionistsTransferrules(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/reception/customer-facing-agents/value/transfer-rules/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/customer-facing-agents/value/transfer-rules/value2")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionReceptionistsTransferrules(t, 422, `{
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

	request := &receptionists.DeleteTransferRulesRequest{
		XiWorkspaceId: "value",
	}
	_, err := client.Delete(context.Background(), string("value"), string("value2"), request)
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

type wireRequestReceptionReceptionistsTransferrules struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionReceptionistsTransferrules records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionReceptionistsTransferrules(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionReceptionistsTransferrules) {
	t.Helper()
	recorded := make(chan wireRequestReceptionReceptionistsTransferrules, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionReceptionistsTransferrules{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionReceptionistsTransferrules(t *testing.T, recorded <-chan wireRequestReceptionReceptionistsTransferrules) wireRequestReceptionReceptionistsTransferrules {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionReceptionistsTransferrules{}
	}
}

// wireJSONContainsReceptionReceptionistsTransferrules reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionReceptionistsTransferrules(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionReceptionistsTransferrules(got, want)
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

// setWireValueReceptionReceptionistsTransferrules sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionReceptionistsTransferrules(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionReceptionistsTransferrules(got, want []byte) bool {
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
