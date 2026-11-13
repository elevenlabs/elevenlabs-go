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

	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestPauseAllWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsClient(t, 200, `{
  "paused_campaign_ids": [
    "value"
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.PauseAll(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsClient(gotResponse, []byte(`{
  "paused_campaign_ids": [
    "value"
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "paused_campaign_ids": [
    "value"
  ]
}`)
	}

	req := receiveWireAgentsOutboundcampaignsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/outbound-campaigns/campaigns/pause-all" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/campaigns/pause-all")
	}
}

func TestPauseAllWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsClient(t, 422, `{
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

	_, err := client.PauseAll(context.Background())
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

func TestStopAllWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsClient(t, 200, `{
  "stopped_campaign_ids": [
    "value"
  ],
  "stopping_campaign_ids": [
    "value"
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.StopAll(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsClient(gotResponse, []byte(`{
  "stopped_campaign_ids": [
    "value"
  ],
  "stopping_campaign_ids": [
    "value"
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "stopped_campaign_ids": [
    "value"
  ],
  "stopping_campaign_ids": [
    "value"
  ]
}`)
	}

	req := receiveWireAgentsOutboundcampaignsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/outbound-campaigns/campaigns/stop-all" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/campaigns/stop-all")
	}
}

func TestStopAllWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsClient(t, 422, `{
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

	_, err := client.StopAll(context.Background())
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

func TestGetProgressWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsClient(t, 200, `{
  "leads": {
    "total": 123,
    "uncalled": 123,
    "leads_needing_attention": 123,
    "active": 123,
    "finished": 123,
    "finished_by_disposition": {
      "key": 123
    }
  },
  "conversations": {
    "scheduled": 123,
    "scheduled_by_type": {
      "key": 123
    },
    "awaiting_permission": 123,
    "in_progress": 123,
    "awaiting_analysis": 123
  },
  "next_call_time_unix": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetProgress(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsClient(gotResponse, []byte(`{
  "leads": {
    "total": 123,
    "uncalled": 123,
    "leads_needing_attention": 123,
    "active": 123,
    "finished": 123,
    "finished_by_disposition": {
      "key": 123
    }
  },
  "conversations": {
    "scheduled": 123,
    "scheduled_by_type": {
      "key": 123
    },
    "awaiting_permission": 123,
    "in_progress": 123,
    "awaiting_analysis": 123
  },
  "next_call_time_unix": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "leads": {
    "total": 123,
    "uncalled": 123,
    "leads_needing_attention": 123,
    "active": 123,
    "finished": 123,
    "finished_by_disposition": {
      "key": 123
    }
  },
  "conversations": {
    "scheduled": 123,
    "scheduled_by_type": {
      "key": 123
    },
    "awaiting_permission": 123,
    "in_progress": 123,
    "awaiting_analysis": 123
  },
  "next_call_time_unix": 123
}`)
	}

	req := receiveWireAgentsOutboundcampaignsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/outbound-campaigns/campaigns/value/progress" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/campaigns/value/progress")
	}
}

func TestGetProgressWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsClient(t, 422, `{
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

	_, err := client.GetProgress(context.Background(), string("value"))
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

func TestGetFollowUpRequestsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsOutboundcampaignsClient(t, 200, `{
  "callbacks": {
    "requested": 123,
    "requested_by_disposition": {
      "key": 123
    },
    "distinct_leads": 123
  },
  "redirects": {
    "requested": 123,
    "requested_by_disposition": {
      "key": 123
    },
    "distinct_leads": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetFollowUpRequests(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsOutboundcampaignsClient(gotResponse, []byte(`{
  "callbacks": {
    "requested": 123,
    "requested_by_disposition": {
      "key": 123
    },
    "distinct_leads": 123
  },
  "redirects": {
    "requested": 123,
    "requested_by_disposition": {
      "key": 123
    },
    "distinct_leads": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "callbacks": {
    "requested": 123,
    "requested_by_disposition": {
      "key": 123
    },
    "distinct_leads": 123
  },
  "redirects": {
    "requested": 123,
    "requested_by_disposition": {
      "key": 123
    },
    "distinct_leads": 123
  }
}`)
	}

	req := receiveWireAgentsOutboundcampaignsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/outbound-campaigns/campaigns/value/follow-up-requests" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/outbound-campaigns/campaigns/value/follow-up-requests")
	}
}

func TestGetFollowUpRequestsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsOutboundcampaignsClient(t, 422, `{
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

	_, err := client.GetFollowUpRequests(context.Background(), string("value"))
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

type wireRequestAgentsOutboundcampaignsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsOutboundcampaignsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsOutboundcampaignsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsOutboundcampaignsClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsOutboundcampaignsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsOutboundcampaignsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsOutboundcampaignsClient(t *testing.T, recorded <-chan wireRequestAgentsOutboundcampaignsClient) wireRequestAgentsOutboundcampaignsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsOutboundcampaignsClient{}
	}
}

// wireJSONContainsAgentsOutboundcampaignsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsOutboundcampaignsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsOutboundcampaignsClient(got, want)
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

// setWireValueAgentsOutboundcampaignsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsOutboundcampaignsClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsOutboundcampaignsClient(got, want []byte) bool {
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
