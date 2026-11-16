package history

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
	types "github.com/elevenlabs/elevenlabs-go/types"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireHistory(t, 200, `{
  "history": [
    {
      "history_item_id": "value",
      "request_id": "value",
      "voice_id": "value",
      "model_id": "value",
      "voice_name": "value",
      "voice_category": "premade",
      "text": "value",
      "date_unix": 123,
      "character_count_change_from": 123,
      "character_count_change_to": 123,
      "content_type": "value",
      "state": {},
      "settings": {
        "key": {}
      },
      "share_link_id": "value",
      "source": "TTS",
      "alignments": {
        "alignment": {
          "characters": [
            "value"
          ],
          "character_start_times_seconds": [
            9.99
          ],
          "character_end_times_seconds": [
            9.99
          ]
        },
        "normalized_alignment": {
          "characters": [
            "value"
          ],
          "character_start_times_seconds": [
            9.99
          ],
          "character_end_times_seconds": [
            9.99
          ]
        }
      },
      "dialogue": [
        {
          "text": "value",
          "voice_id": "value",
          "voice_name": "value"
        }
      ],
      "output_format": "value"
    }
  ],
  "last_history_item_id": "value",
  "has_more": true,
  "scanned_until": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(123)
	startAfterHistoryItemIdValue := string("value")
	voiceIdValue := string("value")
	modelIdValue := string("value")
	dateBeforeUnixValue := int(123)
	dateAfterUnixValue := int(123)
	searchValue := string("value")
	request := &types.ListHistoryRequest{
		PageSize:                &pageSizeValue,
		StartAfterHistoryItemId: &startAfterHistoryItemIdValue,
		VoiceId:                 &voiceIdValue,
		ModelId:                 &modelIdValue,
		DateBeforeUnix:          &dateBeforeUnixValue,
		DateAfterUnix:           &dateAfterUnixValue,
		Search:                  &searchValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualHistory(gotResponse, []byte(`{
  "history": [
    {
      "history_item_id": "value",
      "request_id": "value",
      "voice_id": "value",
      "model_id": "value",
      "voice_name": "value",
      "voice_category": "premade",
      "text": "value",
      "date_unix": 123,
      "character_count_change_from": 123,
      "character_count_change_to": 123,
      "content_type": "value",
      "state": {},
      "settings": {
        "key": {}
      },
      "share_link_id": "value",
      "source": "TTS",
      "alignments": {
        "alignment": {
          "characters": [
            "value"
          ],
          "character_start_times_seconds": [
            9.99
          ],
          "character_end_times_seconds": [
            9.99
          ]
        },
        "normalized_alignment": {
          "characters": [
            "value"
          ],
          "character_start_times_seconds": [
            9.99
          ],
          "character_end_times_seconds": [
            9.99
          ]
        }
      },
      "dialogue": [
        {
          "text": "value",
          "voice_id": "value",
          "voice_name": "value"
        }
      ],
      "output_format": "value"
    }
  ],
  "last_history_item_id": "value",
  "has_more": true,
  "scanned_until": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "history": [
    {
      "history_item_id": "value",
      "request_id": "value",
      "voice_id": "value",
      "model_id": "value",
      "voice_name": "value",
      "voice_category": "premade",
      "text": "value",
      "date_unix": 123,
      "character_count_change_from": 123,
      "character_count_change_to": 123,
      "content_type": "value",
      "state": {},
      "settings": {
        "key": {}
      },
      "share_link_id": "value",
      "source": "TTS",
      "alignments": {
        "alignment": {
          "characters": [
            "value"
          ],
          "character_start_times_seconds": [
            9.99
          ],
          "character_end_times_seconds": [
            9.99
          ]
        },
        "normalized_alignment": {
          "characters": [
            "value"
          ],
          "character_start_times_seconds": [
            9.99
          ],
          "character_end_times_seconds": [
            9.99
          ]
        }
      },
      "dialogue": [
        {
          "text": "value",
          "voice_id": "value",
          "voice_name": "value"
        }
      ],
      "output_format": "value"
    }
  ],
  "last_history_item_id": "value",
  "has_more": true,
  "scanned_until": 123
}`)
	}

	req := receiveWireHistory(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/history" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/history")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
	if got := req.Query.Get("start_after_history_item_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "start_after_history_item_id", got, "value")
	}
	if got := req.Query.Get("voice_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "voice_id", got, "value")
	}
	if got := req.Query.Get("model_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "model_id", got, "value")
	}
	if got := req.Query.Get("date_before_unix"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "date_before_unix", got, "123")
	}
	if got := req.Query.Get("date_after_unix"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "date_after_unix", got, "123")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireHistory(t, 422, `{
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
	startAfterHistoryItemIdValue := string("value")
	voiceIdValue := string("value")
	modelIdValue := string("value")
	dateBeforeUnixValue := int(123)
	dateAfterUnixValue := int(123)
	searchValue := string("value")
	request := &types.ListHistoryRequest{
		PageSize:                &pageSizeValue,
		StartAfterHistoryItemId: &startAfterHistoryItemIdValue,
		VoiceId:                 &voiceIdValue,
		ModelId:                 &modelIdValue,
		DateBeforeUnix:          &dateBeforeUnixValue,
		DateAfterUnix:           &dateAfterUnixValue,
		Search:                  &searchValue,
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireHistory(t, 200, `{
  "history_item_id": "value",
  "request_id": "value",
  "voice_id": "value",
  "model_id": "value",
  "voice_name": "value",
  "voice_category": "premade",
  "text": "value",
  "date_unix": 123,
  "character_count_change_from": 123,
  "character_count_change_to": 123,
  "content_type": "value",
  "state": {},
  "settings": {
    "key": {}
  },
  "share_link_id": "value",
  "source": "TTS",
  "alignments": {
    "alignment": {
      "characters": [
        "value"
      ],
      "character_start_times_seconds": [
        9.99
      ],
      "character_end_times_seconds": [
        9.99
      ]
    },
    "normalized_alignment": {
      "characters": [
        "value"
      ],
      "character_start_times_seconds": [
        9.99
      ],
      "character_end_times_seconds": [
        9.99
      ]
    }
  },
  "dialogue": [
    {
      "text": "value",
      "voice_id": "value",
      "voice_name": "value"
    }
  ],
  "output_format": "value"
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
	if !wireJSONBytesEqualHistory(gotResponse, []byte(`{
  "history_item_id": "value",
  "request_id": "value",
  "voice_id": "value",
  "model_id": "value",
  "voice_name": "value",
  "voice_category": "premade",
  "text": "value",
  "date_unix": 123,
  "character_count_change_from": 123,
  "character_count_change_to": 123,
  "content_type": "value",
  "state": {},
  "settings": {
    "key": {}
  },
  "share_link_id": "value",
  "source": "TTS",
  "alignments": {
    "alignment": {
      "characters": [
        "value"
      ],
      "character_start_times_seconds": [
        9.99
      ],
      "character_end_times_seconds": [
        9.99
      ]
    },
    "normalized_alignment": {
      "characters": [
        "value"
      ],
      "character_start_times_seconds": [
        9.99
      ],
      "character_end_times_seconds": [
        9.99
      ]
    }
  },
  "dialogue": [
    {
      "text": "value",
      "voice_id": "value",
      "voice_name": "value"
    }
  ],
  "output_format": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "history_item_id": "value",
  "request_id": "value",
  "voice_id": "value",
  "model_id": "value",
  "voice_name": "value",
  "voice_category": "premade",
  "text": "value",
  "date_unix": 123,
  "character_count_change_from": 123,
  "character_count_change_to": 123,
  "content_type": "value",
  "state": {},
  "settings": {
    "key": {}
  },
  "share_link_id": "value",
  "source": "TTS",
  "alignments": {
    "alignment": {
      "characters": [
        "value"
      ],
      "character_start_times_seconds": [
        9.99
      ],
      "character_end_times_seconds": [
        9.99
      ]
    },
    "normalized_alignment": {
      "characters": [
        "value"
      ],
      "character_start_times_seconds": [
        9.99
      ],
      "character_end_times_seconds": [
        9.99
      ]
    }
  },
  "dialogue": [
    {
      "text": "value",
      "voice_id": "value",
      "voice_name": "value"
    }
  ],
  "output_format": "value"
}`)
	}

	req := receiveWireHistory(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/history/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/history/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireHistory(t, 422, `{
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

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireHistory(t, 200, `{
  "status": "value"
}`)

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
	if !wireJSONBytesEqualHistory(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireHistory(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/history/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/history/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireHistory(t, 422, `{
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

type wireRequestHistory struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireHistory records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireHistory(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestHistory) {
	t.Helper()
	recorded := make(chan wireRequestHistory, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestHistory{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireHistory(t *testing.T, recorded <-chan wireRequestHistory) wireRequestHistory {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestHistory{}
	}
}

// wireJSONContainsHistory reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsHistory(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualHistory(got, want)
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

// setWireValueHistory sets a named string, int or float field, or the value a pointer field points to.
func setWireValueHistory(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualHistory(got, want []byte) bool {
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
