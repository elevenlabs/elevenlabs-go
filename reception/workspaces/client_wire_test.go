package workspaces

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

	server, recorded := serveWireReceptionWorkspaces(t, 200, `{
  "workspaces": [
    {
      "workspace_id": "value",
      "business_name": "value",
      "billing_tier": "value",
      "receptionist_count": 123,
      "calls_last_30_days": 123,
      "created_at_unix": 123,
      "subscription": {
        "tier_id": "value",
        "billing_period": "monthly_period",
        "billing_period_end_unix": 123,
        "cancel_at_unix": 123,
        "pending_downgrade_tier": "value",
        "pending_downgrade_at_unix": 123
      },
      "has_payment_issue": true,
      "onboarding_completed_at_unix": 123
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
	request := &reception.ListWorkspacesRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionWorkspaces(gotResponse, []byte(`{
  "workspaces": [
    {
      "workspace_id": "value",
      "business_name": "value",
      "billing_tier": "value",
      "receptionist_count": 123,
      "calls_last_30_days": 123,
      "created_at_unix": 123,
      "subscription": {
        "tier_id": "value",
        "billing_period": "monthly_period",
        "billing_period_end_unix": 123,
        "cancel_at_unix": 123,
        "pending_downgrade_tier": "value",
        "pending_downgrade_at_unix": 123
      },
      "has_payment_issue": true,
      "onboarding_completed_at_unix": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "workspaces": [
    {
      "workspace_id": "value",
      "business_name": "value",
      "billing_tier": "value",
      "receptionist_count": 123,
      "calls_last_30_days": 123,
      "created_at_unix": 123,
      "subscription": {
        "tier_id": "value",
        "billing_period": "monthly_period",
        "billing_period_end_unix": 123,
        "cancel_at_unix": 123,
        "pending_downgrade_tier": "value",
        "pending_downgrade_at_unix": 123
      },
      "has_payment_issue": true,
      "onboarding_completed_at_unix": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireReceptionWorkspaces(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/reseller/workspaces" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/workspaces")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionWorkspaces(t, 422, `{
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
	request := &reception.ListWorkspacesRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
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

	server, recorded := serveWireReceptionWorkspaces(t, 200, `{
  "workspace_id": "value",
  "business_name": "value",
  "billing_tier": "value",
  "receptionist_count": 123,
  "calls_last_30_days": 123,
  "created_at_unix": 123,
  "subscription": {
    "tier_id": "value",
    "billing_period": "monthly_period",
    "billing_period_end_unix": 123,
    "cancel_at_unix": 123,
    "pending_downgrade_tier": "value",
    "pending_downgrade_at_unix": 123
  },
  "has_payment_issue": true,
  "onboarding_completed_at_unix": 123,
  "member_results": [
    {
      "email": "value",
      "outcome": "invited",
      "error_status": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ProvisionResellerWorkspaceRequest
	if err := json.Unmarshal([]byte(`{
  "business_name": "value",
  "subscription": {
    "tier_id": "value",
    "billing_period": "monthly_period"
  },
  "phone_number": "value",
  "members": [
    {
      "email": "value",
      "role": "workspace_admin"
    }
  ],
  "complete_onboarding": true,
  "idempotency_key": "value"
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
	if !wireJSONBytesEqualReceptionWorkspaces(gotResponse, []byte(`{
  "workspace_id": "value",
  "business_name": "value",
  "billing_tier": "value",
  "receptionist_count": 123,
  "calls_last_30_days": 123,
  "created_at_unix": 123,
  "subscription": {
    "tier_id": "value",
    "billing_period": "monthly_period",
    "billing_period_end_unix": 123,
    "cancel_at_unix": 123,
    "pending_downgrade_tier": "value",
    "pending_downgrade_at_unix": 123
  },
  "has_payment_issue": true,
  "onboarding_completed_at_unix": 123,
  "member_results": [
    {
      "email": "value",
      "outcome": "invited",
      "error_status": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "workspace_id": "value",
  "business_name": "value",
  "billing_tier": "value",
  "receptionist_count": 123,
  "calls_last_30_days": 123,
  "created_at_unix": 123,
  "subscription": {
    "tier_id": "value",
    "billing_period": "monthly_period",
    "billing_period_end_unix": 123,
    "cancel_at_unix": 123,
    "pending_downgrade_tier": "value",
    "pending_downgrade_at_unix": 123
  },
  "has_payment_issue": true,
  "onboarding_completed_at_unix": 123,
  "member_results": [
    {
      "email": "value",
      "outcome": "invited",
      "error_status": "value"
    }
  ]
}`)
	}

	req := receiveWireReceptionWorkspaces(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/reseller/workspaces" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/workspaces")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionWorkspaces(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionWorkspaces(t, 422, `{
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

	var request reception.ProvisionResellerWorkspaceRequest
	if err := json.Unmarshal([]byte(`{
  "business_name": "value",
  "subscription": {
    "tier_id": "value",
    "billing_period": "monthly_period"
  },
  "phone_number": "value",
  "members": [
    {
      "email": "value",
      "role": "workspace_admin"
    }
  ],
  "complete_onboarding": true,
  "idempotency_key": "value"
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

func TestCompleteOnboardingWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionWorkspaces(t, 200, `{
  "onboarding_completed_at_unix": 123,
  "customer_facing_agent_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.CompleteOnboarding(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionWorkspaces(gotResponse, []byte(`{
  "onboarding_completed_at_unix": 123,
  "customer_facing_agent_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "onboarding_completed_at_unix": 123,
  "customer_facing_agent_id": "value"
}`)
	}

	req := receiveWireReceptionWorkspaces(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/reseller/workspaces/value/complete-onboarding" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/reseller/workspaces/value/complete-onboarding")
	}
}

func TestCompleteOnboardingWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionWorkspaces(t, 422, `{
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

	_, err := client.CompleteOnboarding(context.Background(), string("value"))
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

type wireRequestReceptionWorkspaces struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionWorkspaces records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionWorkspaces(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionWorkspaces) {
	t.Helper()
	recorded := make(chan wireRequestReceptionWorkspaces, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionWorkspaces{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionWorkspaces(t *testing.T, recorded <-chan wireRequestReceptionWorkspaces) wireRequestReceptionWorkspaces {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionWorkspaces{}
	}
}

// wireJSONContainsReceptionWorkspaces reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionWorkspaces(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionWorkspaces(got, want)
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

// setWireValueReceptionWorkspaces sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionWorkspaces(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionWorkspaces(got, want []byte) bool {
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
