package calls

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

	server, recorded := serveWireAgentsOutboundcampaignsCalls(t, 200, `{
  "calls": [
    {
      "intent_id": "value",
      "lead_id": "value",
      "intent_type": "initial",
      "cold_attempt": 123,
      "status": "queued",
      "phone_number": "value",
      "scheduled_time_unix": 123,
      "conversation_id": "value",
      "outcome": "pending",
      "call_result": "answered",
      "outcome_at_unix": 123,
      "decision_disposition": "completed",
      "collected_dynamic_variables": {
        "key": {}
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
	leadIdValue := string("value")
	cursorValue := string("value")
	request := &outboundcampaigns.ListCallsRequest{
		PageSize: &pageSizeValue,
		LeadId:   &leadIdValue,
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
	if !wireJSONBytesEqualAgentsOutboundcampaignsCalls(gotResponse, []byte(`{
  "calls": [
    {
      "intent_id": "value",
      "lead_id": "value",
      "intent_type": "initial",
      "cold_attempt": 123,
      "status": "queued",
      "phone_number": "value",
      "scheduled_time_unix": 123,
      "conversation_id": "value",
      "outcome": "pending",
      "call_result": "answered",
      "outcome_at_unix": 123,
      "decision_disposition": "completed",
      "collected_dynamic_variables": {
        "key": {}
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
  "calls": [
    {
      "intent_id": "value",
      "lead_id": "value",
      "intent_type": "initial",
      "cold_attempt": 123,
      "status": "queued",
      "phone_number": "value",
      "scheduled_time_unix": 123,
      "conversation_id": "value",
      "outcome": "pending",
      "call_result": "answered",
      "outcome_at_unix": 123,
      "decision_disposition": "completed",
      "collected_dynamic_variables": {
        "key": {}
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

	req := receiveWireAgentsOutboundcampaignsCalls(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/outbound-campaigns/campaigns/value/calls" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/campaigns/value/calls")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
	if got := req.Query.Get("lead_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "lead_id", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsCalls(t, 422, `{
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
	leadIdValue := string("value")
	cursorValue := string("value")
	request := &outboundcampaigns.ListCallsRequest{
		PageSize: &pageSizeValue,
		LeadId:   &leadIdValue,
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

type wireRequestAgentsOutboundcampaignsCalls struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsOutboundcampaignsCalls records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsOutboundcampaignsCalls(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsOutboundcampaignsCalls) {
	t.Helper()
	recorded := make(chan wireRequestAgentsOutboundcampaignsCalls, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsOutboundcampaignsCalls{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsOutboundcampaignsCalls(t *testing.T, recorded <-chan wireRequestAgentsOutboundcampaignsCalls) wireRequestAgentsOutboundcampaignsCalls {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsOutboundcampaignsCalls{}
	}
}

// wireJSONContainsAgentsOutboundcampaignsCalls reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsOutboundcampaignsCalls(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsOutboundcampaignsCalls(got, want)
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

// setWireValueAgentsOutboundcampaignsCalls sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsOutboundcampaignsCalls(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsOutboundcampaignsCalls(got, want []byte) bool {
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
