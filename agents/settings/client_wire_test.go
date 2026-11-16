package settings

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

	server, recorded := serveWireAgentsSettings(t, 200, `{
  "conversation_initiation_client_data_webhook": {
    "url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "webhooks": {
    "post_call_webhook_id": "value",
    "events": [
      "transcript"
    ],
    "transcript_format": "json",
    "exclude_transcript": true,
    "send_audio": true
  },
  "can_use_mcp_servers": true,
  "rag_retention_period_days": 30,
  "conversation_embedding_retention_days": 123,
  "default_livekit_stack": "standard"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsSettings(gotResponse, []byte(`{
  "conversation_initiation_client_data_webhook": {
    "url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "webhooks": {
    "post_call_webhook_id": "value",
    "events": [
      "transcript"
    ],
    "transcript_format": "json",
    "exclude_transcript": true,
    "send_audio": true
  },
  "can_use_mcp_servers": true,
  "rag_retention_period_days": 30,
  "conversation_embedding_retention_days": 123,
  "default_livekit_stack": "standard"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "conversation_initiation_client_data_webhook": {
    "url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "webhooks": {
    "post_call_webhook_id": "value",
    "events": [
      "transcript"
    ],
    "transcript_format": "json",
    "exclude_transcript": true,
    "send_audio": true
  },
  "can_use_mcp_servers": true,
  "rag_retention_period_days": 30,
  "conversation_embedding_retention_days": 123,
  "default_livekit_stack": "standard"
}`)
	}

	req := receiveWireAgentsSettings(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/settings" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/settings")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsSettings(t, 422, `{
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

	_, err := client.Get(context.Background())
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

	server, recorded := serveWireAgentsSettings(t, 200, `{
  "conversation_initiation_client_data_webhook": {
    "url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "webhooks": {
    "post_call_webhook_id": "value",
    "events": [
      "transcript"
    ],
    "transcript_format": "json",
    "exclude_transcript": true,
    "send_audio": true
  },
  "can_use_mcp_servers": true,
  "rag_retention_period_days": 30,
  "conversation_embedding_retention_days": 123,
  "default_livekit_stack": "standard"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.PatchConvAiSettingsRequest
	if err := json.Unmarshal([]byte(`{
  "conversation_initiation_client_data_webhook": {
    "url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "webhooks": {
    "post_call_webhook_id": "value",
    "events": [
      "transcript"
    ],
    "transcript_format": "json",
    "exclude_transcript": true,
    "send_audio": true
  },
  "can_use_mcp_servers": true,
  "rag_retention_period_days": 30,
  "conversation_embedding_retention_days": 123,
  "default_livekit_stack": "standard"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsSettings(gotResponse, []byte(`{
  "conversation_initiation_client_data_webhook": {
    "url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "webhooks": {
    "post_call_webhook_id": "value",
    "events": [
      "transcript"
    ],
    "transcript_format": "json",
    "exclude_transcript": true,
    "send_audio": true
  },
  "can_use_mcp_servers": true,
  "rag_retention_period_days": 30,
  "conversation_embedding_retention_days": 123,
  "default_livekit_stack": "standard"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "conversation_initiation_client_data_webhook": {
    "url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "webhooks": {
    "post_call_webhook_id": "value",
    "events": [
      "transcript"
    ],
    "transcript_format": "json",
    "exclude_transcript": true,
    "send_audio": true
  },
  "can_use_mcp_servers": true,
  "rag_retention_period_days": 30,
  "conversation_embedding_retention_days": 123,
  "default_livekit_stack": "standard"
}`)
	}

	req := receiveWireAgentsSettings(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/settings" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/settings")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsSettings(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsSettings(t, 422, `{
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

	var request agents.PatchConvAiSettingsRequest
	if err := json.Unmarshal([]byte(`{
  "conversation_initiation_client_data_webhook": {
    "url": "value",
    "request_headers": {
      "key": "value"
    }
  },
  "webhooks": {
    "post_call_webhook_id": "value",
    "events": [
      "transcript"
    ],
    "transcript_format": "json",
    "exclude_transcript": true,
    "send_audio": true
  },
  "can_use_mcp_servers": true,
  "rag_retention_period_days": 30,
  "conversation_embedding_retention_days": 123,
  "default_livekit_stack": "standard"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), &request)
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

type wireRequestAgentsSettings struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsSettings records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsSettings(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsSettings) {
	t.Helper()
	recorded := make(chan wireRequestAgentsSettings, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsSettings{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsSettings(t *testing.T, recorded <-chan wireRequestAgentsSettings) wireRequestAgentsSettings {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsSettings{}
	}
}

// wireJSONContainsAgentsSettings reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsSettings(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsSettings(got, want)
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

// setWireValueAgentsSettings sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsSettings(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsSettings(got, want []byte) bool {
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
