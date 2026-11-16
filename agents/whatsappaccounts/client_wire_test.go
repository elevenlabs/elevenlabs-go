package whatsappaccounts

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

	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsWhatsappaccounts(t, 200, `{
  "business_account_id": "value",
  "phone_number_id": "value",
  "business_account_name": "value",
  "phone_number_name": "value",
  "phone_number": "value",
  "account_type": "cloud_api",
  "assigned_agent_id": "value",
  "enable_messaging": true,
  "enable_audio_message_response": true,
  "enable_typing_indicator": true,
  "assigned_agent_name": "value",
  "is_token_expired": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsWhatsappaccounts(gotResponse, []byte(`{
  "business_account_id": "value",
  "phone_number_id": "value",
  "business_account_name": "value",
  "phone_number_name": "value",
  "phone_number": "value",
  "account_type": "cloud_api",
  "assigned_agent_id": "value",
  "enable_messaging": true,
  "enable_audio_message_response": true,
  "enable_typing_indicator": true,
  "assigned_agent_name": "value",
  "is_token_expired": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "business_account_id": "value",
  "phone_number_id": "value",
  "business_account_name": "value",
  "phone_number_name": "value",
  "phone_number": "value",
  "account_type": "cloud_api",
  "assigned_agent_id": "value",
  "enable_messaging": true,
  "enable_audio_message_response": true,
  "enable_typing_indicator": true,
  "assigned_agent_name": "value",
  "is_token_expired": true
}`)
	}

	req := receiveWireAgentsWhatsappaccounts(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/whatsapp-accounts/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/whatsapp-accounts/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsWhatsappaccounts(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"))
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

	server, recorded := serveWireAgentsWhatsappaccounts(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.UpdateWhatsAppAccountRequest
	if err := json.Unmarshal([]byte(`{
  "phone_number_id": "value",
  "assigned_agent_id": "value",
  "enable_messaging": true,
  "enable_audio_message_response": true,
  "enable_typing_indicator": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsWhatsappaccounts(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsWhatsappaccounts(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/whatsapp-accounts/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/whatsapp-accounts/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsWhatsappaccounts(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsWhatsappaccounts(t, 422, `{
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

	var request agents.UpdateWhatsAppAccountRequest
	if err := json.Unmarshal([]byte(`{
  "phone_number_id": "value",
  "assigned_agent_id": "value",
  "enable_messaging": true,
  "enable_audio_message_response": true,
  "enable_typing_indicator": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), &request)
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

	server, recorded := serveWireAgentsWhatsappaccounts(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsWhatsappaccounts(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsWhatsappaccounts(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/whatsapp-accounts/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/whatsapp-accounts/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsWhatsappaccounts(t, 422, `{
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

	_, err := client.Delete(context.Background(), string("value"))
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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsWhatsappaccounts(t, 200, `{
  "items": [
    {
      "business_account_id": "value",
      "phone_number_id": "value",
      "business_account_name": "value",
      "phone_number_name": "value",
      "phone_number": "value",
      "account_type": "cloud_api",
      "assigned_agent_id": "value",
      "enable_messaging": true,
      "enable_audio_message_response": true,
      "enable_typing_indicator": true,
      "assigned_agent_name": "value",
      "is_token_expired": true
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	agentIdValue := string("value")
	request := &agents.ListWhatsappAccountsRequest{
		AgentId: &agentIdValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsWhatsappaccounts(gotResponse, []byte(`{
  "items": [
    {
      "business_account_id": "value",
      "phone_number_id": "value",
      "business_account_name": "value",
      "phone_number_name": "value",
      "phone_number": "value",
      "account_type": "cloud_api",
      "assigned_agent_id": "value",
      "enable_messaging": true,
      "enable_audio_message_response": true,
      "enable_typing_indicator": true,
      "assigned_agent_name": "value",
      "is_token_expired": true
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "items": [
    {
      "business_account_id": "value",
      "phone_number_id": "value",
      "business_account_name": "value",
      "phone_number_name": "value",
      "phone_number": "value",
      "account_type": "cloud_api",
      "assigned_agent_id": "value",
      "enable_messaging": true,
      "enable_audio_message_response": true,
      "enable_typing_indicator": true,
      "assigned_agent_name": "value",
      "is_token_expired": true
    }
  ]
}`)
	}

	req := receiveWireAgentsWhatsappaccounts(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/whatsapp-accounts" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/whatsapp-accounts")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsWhatsappaccounts(t, 422, `{
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

	agentIdValue := string("value")
	request := &agents.ListWhatsappAccountsRequest{
		AgentId: &agentIdValue,
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

type wireRequestAgentsWhatsappaccounts struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsWhatsappaccounts records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsWhatsappaccounts(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsWhatsappaccounts) {
	t.Helper()
	recorded := make(chan wireRequestAgentsWhatsappaccounts, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsWhatsappaccounts{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsWhatsappaccounts(t *testing.T, recorded <-chan wireRequestAgentsWhatsappaccounts) wireRequestAgentsWhatsappaccounts {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsWhatsappaccounts{}
	}
}

// wireJSONContainsAgentsWhatsappaccounts reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsWhatsappaccounts(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsWhatsappaccounts(got, want)
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

// setWireValueAgentsWhatsappaccounts sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsWhatsappaccounts(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsWhatsappaccounts(got, want []byte) bool {
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
