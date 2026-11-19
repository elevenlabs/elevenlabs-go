package topics

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

	conversations "github.com/elevenlabs/elevenlabs-go/agents/conversations"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsConversationsTopics(t, 200, `{
  "topics": [
    {
      "topic_id": "value",
      "label": "value",
      "description": "value",
      "conversation_count": 123,
      "parent_topic_id": "value",
      "x_2d": 9.99,
      "y_2d": 9.99,
      "metrics": {
        "conversation_count": 123,
        "sentiment": {
          "sentiment": {
            "count": 123,
            "sum": 9.99,
            "min": 9.99,
            "max": 9.99
          },
          "frustration": {
            "count": 123,
            "sum": 9.99,
            "min": 9.99,
            "max": 9.99
          },
          "positive_count": 123,
          "neutral_count": 123,
          "negative_count": 123
        },
        "evaluation_criteria": [
          {
            "criteria_id": "value",
            "success_count": 123,
            "failure_count": 123,
            "unknown_count": 123
          }
        ]
      },
      "success_rate": 9.99
    }
  ],
  "window_start_unix_secs": 123,
  "window_end_unix_secs": 123,
  "aggregated_run_count": 123,
  "has_more": true,
  "next_cursor": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	fromUnixSecsValue := int(123)
	toUnixSecsValue := int(123)
	includeEvaluationCriteriaValue := bool(true)
	cursorValue := string("value")
	request := &conversations.GetTopicsRequest{
		PageSize:                  &pageSizeValue,
		FromUnixSecs:              &fromUnixSecsValue,
		ToUnixSecs:                &toUnixSecsValue,
		IncludeEvaluationCriteria: &includeEvaluationCriteriaValue,
		Cursor:                    &cursorValue,
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsConversationsTopics(gotResponse, []byte(`{
  "topics": [
    {
      "topic_id": "value",
      "label": "value",
      "description": "value",
      "conversation_count": 123,
      "parent_topic_id": "value",
      "x_2d": 9.99,
      "y_2d": 9.99,
      "metrics": {
        "conversation_count": 123,
        "sentiment": {
          "sentiment": {
            "count": 123,
            "sum": 9.99,
            "min": 9.99,
            "max": 9.99
          },
          "frustration": {
            "count": 123,
            "sum": 9.99,
            "min": 9.99,
            "max": 9.99
          },
          "positive_count": 123,
          "neutral_count": 123,
          "negative_count": 123
        },
        "evaluation_criteria": [
          {
            "criteria_id": "value",
            "success_count": 123,
            "failure_count": 123,
            "unknown_count": 123
          }
        ]
      },
      "success_rate": 9.99
    }
  ],
  "window_start_unix_secs": 123,
  "window_end_unix_secs": 123,
  "aggregated_run_count": 123,
  "has_more": true,
  "next_cursor": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "topics": [
    {
      "topic_id": "value",
      "label": "value",
      "description": "value",
      "conversation_count": 123,
      "parent_topic_id": "value",
      "x_2d": 9.99,
      "y_2d": 9.99,
      "metrics": {
        "conversation_count": 123,
        "sentiment": {
          "sentiment": {
            "count": 123,
            "sum": 9.99,
            "min": 9.99,
            "max": 9.99
          },
          "frustration": {
            "count": 123,
            "sum": 9.99,
            "min": 9.99,
            "max": 9.99
          },
          "positive_count": 123,
          "neutral_count": 123,
          "negative_count": 123
        },
        "evaluation_criteria": [
          {
            "criteria_id": "value",
            "success_count": 123,
            "failure_count": 123,
            "unknown_count": 123
          }
        ]
      },
      "success_rate": 9.99
    }
  ],
  "window_start_unix_secs": 123,
  "window_end_unix_secs": 123,
  "aggregated_run_count": 123,
  "has_more": true,
  "next_cursor": "value"
}`)
	}

	req := receiveWireAgentsConversationsTopics(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agents/value/topics" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/topics")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("from_unix_secs"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "from_unix_secs", got, "123")
	}
	if got := req.Query.Get("to_unix_secs"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "to_unix_secs", got, "123")
	}
	if got := req.Query.Get("include_evaluation_criteria"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "include_evaluation_criteria", got, "true")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsConversationsTopics(t, 422, `{
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
	fromUnixSecsValue := int(123)
	toUnixSecsValue := int(123)
	includeEvaluationCriteriaValue := bool(true)
	cursorValue := string("value")
	request := &conversations.GetTopicsRequest{
		PageSize:                  &pageSizeValue,
		FromUnixSecs:              &fromUnixSecsValue,
		ToUnixSecs:                &toUnixSecsValue,
		IncludeEvaluationCriteria: &includeEvaluationCriteriaValue,
		Cursor:                    &cursorValue,
	}
	_, err := client.Get(context.Background(), string("value"), request)
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

type wireRequestAgentsConversationsTopics struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsConversationsTopics records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsConversationsTopics(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsConversationsTopics) {
	t.Helper()
	recorded := make(chan wireRequestAgentsConversationsTopics, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsConversationsTopics{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsConversationsTopics(t *testing.T, recorded <-chan wireRequestAgentsConversationsTopics) wireRequestAgentsConversationsTopics {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsConversationsTopics{}
	}
}

// wireJSONContainsAgentsConversationsTopics reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsConversationsTopics(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsConversationsTopics(got, want)
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

// setWireValueAgentsConversationsTopics sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsConversationsTopics(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsConversationsTopics(got, want []byte) bool {
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
