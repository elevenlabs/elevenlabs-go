package phonenumbers

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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionPhonenumbers(t, 200, `[
  {
    "phone_number": "value",
    "is_byo": true,
    "provider": "twilio",
    "nanp_area_code": 123,
    "customer_facing_agent_id": "value",
    "deactivated_at_unix": 123,
    "deactivation_reason": "tier_downgrade",
    "regulatory_status": "held",
    "regulatory_number_type": "local",
    "regulatory_country": "value",
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "is_active": true,
    "accepts_public_calls": true
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.ListPhoneNumbersRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionPhonenumbers(gotResponse, []byte(`[
  {
    "phone_number": "value",
    "is_byo": true,
    "provider": "twilio",
    "nanp_area_code": 123,
    "customer_facing_agent_id": "value",
    "deactivated_at_unix": 123,
    "deactivation_reason": "tier_downgrade",
    "regulatory_status": "held",
    "regulatory_number_type": "local",
    "regulatory_country": "value",
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "is_active": true,
    "accepts_public_calls": true
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "phone_number": "value",
    "is_byo": true,
    "provider": "twilio",
    "nanp_area_code": 123,
    "customer_facing_agent_id": "value",
    "deactivated_at_unix": 123,
    "deactivation_reason": "tier_downgrade",
    "regulatory_status": "held",
    "regulatory_number_type": "local",
    "regulatory_country": "value",
    "id": "value",
    "created_at_unix": 123,
    "updated_at_unix": 123,
    "is_active": true,
    "accepts_public_calls": true
  }
]`)
	}

	req := receiveWireReceptionPhonenumbers(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/phone-numbers/assignments" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/phone-numbers/assignments")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionPhonenumbers(t, 422, `{
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

	request := &reception.ListPhoneNumbersRequest{
		XiWorkspaceId: "value",
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

	server, recorded := serveWireReceptionPhonenumbers(t, 201, `{
  "phone_number": "value",
  "is_byo": true,
  "provider": "twilio",
  "nanp_area_code": 123,
  "customer_facing_agent_id": "value",
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "regulatory_status": "held",
  "regulatory_number_type": "local",
  "regulatory_country": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "is_active": true,
  "accepts_public_calls": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.CreatePhoneAssignmentRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "customer_facing_agent_id": "value",
  "area_codes": [
    "value"
  ],
  "number_type": "local",
  "idempotency_key": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionPhonenumbers(gotResponse, []byte(`{
  "phone_number": "value",
  "is_byo": true,
  "provider": "twilio",
  "nanp_area_code": 123,
  "customer_facing_agent_id": "value",
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "regulatory_status": "held",
  "regulatory_number_type": "local",
  "regulatory_country": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "is_active": true,
  "accepts_public_calls": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "phone_number": "value",
  "is_byo": true,
  "provider": "twilio",
  "nanp_area_code": 123,
  "customer_facing_agent_id": "value",
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "regulatory_status": "held",
  "regulatory_number_type": "local",
  "regulatory_country": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "is_active": true,
  "accepts_public_calls": true
}`)
	}

	req := receiveWireReceptionPhonenumbers(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/phone-numbers/assignments" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/phone-numbers/assignments")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionPhonenumbers(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionPhonenumbers(t, 422, `{
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

	var request reception.CreatePhoneAssignmentRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "customer_facing_agent_id": "value",
  "area_codes": [
    "value"
  ],
  "number_type": "local",
  "idempotency_key": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionPhonenumbers(t, 200, `{
  "phone_number": "value",
  "is_byo": true,
  "provider": "twilio",
  "nanp_area_code": 123,
  "customer_facing_agent_id": "value",
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "regulatory_status": "held",
  "regulatory_number_type": "local",
  "regulatory_country": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "is_active": true,
  "accepts_public_calls": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReassignPhoneRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "phone_assignment_id": "value",
  "customer_facing_agent_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionPhonenumbers(gotResponse, []byte(`{
  "phone_number": "value",
  "is_byo": true,
  "provider": "twilio",
  "nanp_area_code": 123,
  "customer_facing_agent_id": "value",
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "regulatory_status": "held",
  "regulatory_number_type": "local",
  "regulatory_country": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "is_active": true,
  "accepts_public_calls": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "phone_number": "value",
  "is_byo": true,
  "provider": "twilio",
  "nanp_area_code": 123,
  "customer_facing_agent_id": "value",
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "regulatory_status": "held",
  "regulatory_number_type": "local",
  "regulatory_country": "value",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "is_active": true,
  "accepts_public_calls": true
}`)
	}

	req := receiveWireReceptionPhonenumbers(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/reception/phone-numbers/assignments/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/phone-numbers/assignments/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionPhonenumbers(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionPhonenumbers(t, 422, `{
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

	var request reception.ReassignPhoneRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "phone_assignment_id": "value",
  "customer_facing_agent_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
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

	server, recorded := serveWireReceptionPhonenumbers(t, 204, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.DeletePhoneNumbersRequest{
		XiWorkspaceId: "value",
	}
	err := client.Delete(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireReceptionPhonenumbers(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/reception/phone-numbers/assignments/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/phone-numbers/assignments/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionPhonenumbers(t, 422, `{
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

	request := &reception.DeletePhoneNumbersRequest{
		XiWorkspaceId: "value",
	}
	err := client.Delete(context.Background(), string("value"), request)
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

type wireRequestReceptionPhonenumbers struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionPhonenumbers records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionPhonenumbers(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionPhonenumbers) {
	t.Helper()
	recorded := make(chan wireRequestReceptionPhonenumbers, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionPhonenumbers{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionPhonenumbers(t *testing.T, recorded <-chan wireRequestReceptionPhonenumbers) wireRequestReceptionPhonenumbers {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionPhonenumbers{}
	}
}

// wireJSONContainsReceptionPhonenumbers reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionPhonenumbers(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionPhonenumbers(got, want)
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

// setWireValueReceptionPhonenumbers sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionPhonenumbers(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionPhonenumbers(got, want []byte) bool {
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
