package models

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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireModels(t, 200, `[
  {
    "model_id": "value",
    "name": "value",
    "can_be_finetuned": true,
    "can_do_text_to_speech": true,
    "can_do_voice_conversion": true,
    "can_use_style": true,
    "can_use_speaker_boost": true,
    "serves_pro_voices": true,
    "token_cost_factor": 9.99,
    "description": "value",
    "requires_alpha_access": true,
    "max_characters_request_free_user": 123,
    "max_characters_request_subscribed_user": 123,
    "maximum_text_length_per_request": 123,
    "languages": [
      {
        "language_id": "value",
        "name": "value"
      }
    ],
    "model_rates": {
      "character_cost_multiplier": 9.99,
      "cost_discount_multiplier": 9.99
    },
    "concurrency_group": "value"
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.List(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualModels(gotResponse, []byte(`[
  {
    "model_id": "value",
    "name": "value",
    "can_be_finetuned": true,
    "can_do_text_to_speech": true,
    "can_do_voice_conversion": true,
    "can_use_style": true,
    "can_use_speaker_boost": true,
    "serves_pro_voices": true,
    "token_cost_factor": 9.99,
    "description": "value",
    "requires_alpha_access": true,
    "max_characters_request_free_user": 123,
    "max_characters_request_subscribed_user": 123,
    "maximum_text_length_per_request": 123,
    "languages": [
      {
        "language_id": "value",
        "name": "value"
      }
    ],
    "model_rates": {
      "character_cost_multiplier": 9.99,
      "cost_discount_multiplier": 9.99
    },
    "concurrency_group": "value"
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "model_id": "value",
    "name": "value",
    "can_be_finetuned": true,
    "can_do_text_to_speech": true,
    "can_do_voice_conversion": true,
    "can_use_style": true,
    "can_use_speaker_boost": true,
    "serves_pro_voices": true,
    "token_cost_factor": 9.99,
    "description": "value",
    "requires_alpha_access": true,
    "max_characters_request_free_user": 123,
    "max_characters_request_subscribed_user": 123,
    "maximum_text_length_per_request": 123,
    "languages": [
      {
        "language_id": "value",
        "name": "value"
      }
    ],
    "model_rates": {
      "character_cost_multiplier": 9.99,
      "cost_discount_multiplier": 9.99
    },
    "concurrency_group": "value"
  }
]`)
	}

	req := receiveWireModels(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/models" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/models")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireModels(t, 422, `{
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

	_, err := client.List(context.Background())
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

type wireRequestModels struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireModels records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireModels(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestModels) {
	t.Helper()
	recorded := make(chan wireRequestModels, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestModels{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireModels(t *testing.T, recorded <-chan wireRequestModels) wireRequestModels {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestModels{}
	}
}

// wireJSONContainsModels reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsModels(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualModels(got, want)
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

// setWireValueModels sets a named string, int or float field, or the value a pointer field points to.
func setWireValueModels(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualModels(got, want []byte) bool {
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
