package leads

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

	outboundcampaigns "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsLeads(t, 200, `{
  "leads": [
    {
      "lead_id": "value",
      "recipient_id": "value",
      "phone_number": "value",
      "phone_numbers": [
        "value"
      ],
      "phone_numbers_index": 123,
      "timezone_override": "value",
      "user_id": "value",
      "attempt_count": 123,
      "failure_count": 123,
      "callback_count": 123,
      "redirect_count": 123,
      "do_not_call": true,
      "terminal_disposition": "campaign_ended",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "dynamic_variables": {
        "key": {}
      },
      "overrides": {
        "language": "value",
        "first_message": "value",
        "prompt": "value",
        "voice_id": "value"
      },
      "delivery": {
        "status": "pending",
        "stage": "queued",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_message": "value"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(123)
	searchValue := string("value")
	cursorValue := string("value")
	request := &outboundcampaigns.ListLeadsRequest{
		PageSize: &pageSizeValue,
		Search:   &searchValue,
		Cursor:   &cursorValue,
	}
	result, err := client.List(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsLeads(gotResponse, []byte(`{
  "leads": [
    {
      "lead_id": "value",
      "recipient_id": "value",
      "phone_number": "value",
      "phone_numbers": [
        "value"
      ],
      "phone_numbers_index": 123,
      "timezone_override": "value",
      "user_id": "value",
      "attempt_count": 123,
      "failure_count": 123,
      "callback_count": 123,
      "redirect_count": 123,
      "do_not_call": true,
      "terminal_disposition": "campaign_ended",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "dynamic_variables": {
        "key": {}
      },
      "overrides": {
        "language": "value",
        "first_message": "value",
        "prompt": "value",
        "voice_id": "value"
      },
      "delivery": {
        "status": "pending",
        "stage": "queued",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_message": "value"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "leads": [
    {
      "lead_id": "value",
      "recipient_id": "value",
      "phone_number": "value",
      "phone_numbers": [
        "value"
      ],
      "phone_numbers_index": 123,
      "timezone_override": "value",
      "user_id": "value",
      "attempt_count": 123,
      "failure_count": 123,
      "callback_count": 123,
      "redirect_count": 123,
      "do_not_call": true,
      "terminal_disposition": "campaign_ended",
      "created_at_unix": 123,
      "updated_at_unix": 123,
      "dynamic_variables": {
        "key": {}
      },
      "overrides": {
        "language": "value",
        "first_message": "value",
        "prompt": "value",
        "voice_id": "value"
      },
      "delivery": {
        "status": "pending",
        "stage": "queued",
        "sent_at_unix": 123,
        "delivered_at_unix": 123,
        "read_at_unix": 123,
        "replied_at_unix": 123,
        "failed_at_unix": 123,
        "error_message": "value"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsOutboundcampaignsLeads(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/outbound-campaigns/campaigns/value/leads" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/campaigns/value/leads")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsLeads(t, 422, `{
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

	pageSizeValue := int(123)
	searchValue := string("value")
	cursorValue := string("value")
	request := &outboundcampaigns.ListLeadsRequest{
		PageSize: &pageSizeValue,
		Search:   &searchValue,
		Cursor:   &cursorValue,
	}
	_, err := client.List(context.Background(), string("value"), request)
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsLeads(t, 200, `{
  "lead_id": "value",
  "recipient_id": "value",
  "phone_number": "value",
  "phone_numbers": [
    "value"
  ],
  "phone_numbers_index": 123,
  "timezone_override": "value",
  "user_id": "value",
  "attempt_count": 123,
  "failure_count": 123,
  "callback_count": 123,
  "redirect_count": 123,
  "do_not_call": true,
  "terminal_disposition": "campaign_ended",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "dynamic_variables": {
    "key": {}
  },
  "overrides": {
    "language": "value",
    "first_message": "value",
    "prompt": "value",
    "voice_id": "value"
  },
  "delivery": {
    "status": "pending",
    "stage": "queued",
    "sent_at_unix": 123,
    "delivered_at_unix": 123,
    "read_at_unix": 123,
    "replied_at_unix": 123,
    "failed_at_unix": 123,
    "error_message": "value"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsLeads(gotResponse, []byte(`{
  "lead_id": "value",
  "recipient_id": "value",
  "phone_number": "value",
  "phone_numbers": [
    "value"
  ],
  "phone_numbers_index": 123,
  "timezone_override": "value",
  "user_id": "value",
  "attempt_count": 123,
  "failure_count": 123,
  "callback_count": 123,
  "redirect_count": 123,
  "do_not_call": true,
  "terminal_disposition": "campaign_ended",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "dynamic_variables": {
    "key": {}
  },
  "overrides": {
    "language": "value",
    "first_message": "value",
    "prompt": "value",
    "voice_id": "value"
  },
  "delivery": {
    "status": "pending",
    "stage": "queued",
    "sent_at_unix": 123,
    "delivered_at_unix": 123,
    "read_at_unix": 123,
    "replied_at_unix": 123,
    "failed_at_unix": 123,
    "error_message": "value"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "lead_id": "value",
  "recipient_id": "value",
  "phone_number": "value",
  "phone_numbers": [
    "value"
  ],
  "phone_numbers_index": 123,
  "timezone_override": "value",
  "user_id": "value",
  "attempt_count": 123,
  "failure_count": 123,
  "callback_count": 123,
  "redirect_count": 123,
  "do_not_call": true,
  "terminal_disposition": "campaign_ended",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "dynamic_variables": {
    "key": {}
  },
  "overrides": {
    "language": "value",
    "first_message": "value",
    "prompt": "value",
    "voice_id": "value"
  },
  "delivery": {
    "status": "pending",
    "stage": "queued",
    "sent_at_unix": 123,
    "delivered_at_unix": 123,
    "read_at_unix": 123,
    "replied_at_unix": 123,
    "failed_at_unix": 123,
    "error_message": "value"
  }
}`)
	}

	req := receiveWireAgentsOutboundcampaignsLeads(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/outbound-campaigns/campaigns/value/leads/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/campaigns/value/leads/value2")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsLeads(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"), string("value2"))
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

	server, recorded := serveWireAgentsOutboundcampaignsLeads(t, 200, `{
  "campaign_id": "value",
  "lead_id": "value",
  "has_calls_in_flight": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsLeads(gotResponse, []byte(`{
  "campaign_id": "value",
  "lead_id": "value",
  "has_calls_in_flight": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "campaign_id": "value",
  "lead_id": "value",
  "has_calls_in_flight": true
}`)
	}

	req := receiveWireAgentsOutboundcampaignsLeads(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/outbound-campaigns/campaigns/value/leads/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/campaigns/value/leads/value2")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsLeads(t, 422, `{
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

	_, err := client.Delete(context.Background(), string("value"), string("value2"))
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

type wireRequestAgentsOutboundcampaignsLeads struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsOutboundcampaignsLeads records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsOutboundcampaignsLeads(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsOutboundcampaignsLeads) {
	t.Helper()
	recorded := make(chan wireRequestAgentsOutboundcampaignsLeads, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsOutboundcampaignsLeads{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsOutboundcampaignsLeads(t *testing.T, recorded <-chan wireRequestAgentsOutboundcampaignsLeads) wireRequestAgentsOutboundcampaignsLeads {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsOutboundcampaignsLeads{}
	}
}

// wireJSONContainsAgentsOutboundcampaignsLeads reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsOutboundcampaignsLeads(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsOutboundcampaignsLeads(got, want)
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

// setWireValueAgentsOutboundcampaignsLeads sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsOutboundcampaignsLeads(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsOutboundcampaignsLeads(got, want []byte) bool {
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
