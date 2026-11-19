package texttodialogue

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

func TestFullWithTimestampsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireTexttodialogue(t, 200, `{
  "audio_base64": "value",
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
  },
  "voice_segments": [
    {
      "voice_id": "value",
      "start_time_seconds": 9.99,
      "end_time_seconds": 9.99,
      "character_start_index": 123,
      "character_end_index": 123,
      "dialogue_input_index": 123
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodyTextToDialogueFullWithTimestamps
	if err := json.Unmarshal([]byte(`{
  "enable_logging": true,
  "output_format": "value",
  "inputs": [
    {
      "text": "value",
      "voice_id": "value"
    }
  ],
  "model_id": "value",
  "language_code": "value",
  "settings": {
    "stability": 0.5,
    "similarity": 0.5,
    "clone_strength": 0.5
  },
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "seed": 123,
  "apply_text_normalization": "auto",
  "use_pvc_as_ivc": true,
  "previous_request_ids": [
    "value"
  ],
  "next_request_ids": [
    "value"
  ],
  "previous_text": "value",
  "future_text": "value",
  "audio_effects": {
    "filter_preset_id": "value",
    "distance": 0.5,
    "environment_id": "value",
    "background_noise_id": "value",
    "send_level": 0.5,
    "pan": 0.5
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	outputFormatValue := string("value")
	request.OutputFormat = &outputFormatValue
	enableLoggingValue := bool(true)
	request.EnableLogging = &enableLoggingValue
	result, err := client.FullWithTimestamps(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualTexttodialogue(gotResponse, []byte(`{
  "audio_base64": "value",
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
  },
  "voice_segments": [
    {
      "voice_id": "value",
      "start_time_seconds": 9.99,
      "end_time_seconds": 9.99,
      "character_start_index": 123,
      "character_end_index": 123,
      "dialogue_input_index": 123
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "audio_base64": "value",
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
  },
  "voice_segments": [
    {
      "voice_id": "value",
      "start_time_seconds": 9.99,
      "end_time_seconds": 9.99,
      "character_start_index": 123,
      "character_end_index": 123,
      "dialogue_input_index": 123
    }
  ]
}`)
	}

	req := receiveWireTexttodialogue(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/text-to-dialogue/with-timestamps" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/text-to-dialogue/with-timestamps")
	}
	if got := req.Query.Get("output_format"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "output_format", got, "value")
	}
	if got := req.Query.Get("enable_logging"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "enable_logging", got, "true")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsTexttodialogue(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestFullWithTimestampsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireTexttodialogue(t, 422, `{
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

	var request types.BodyTextToDialogueFullWithTimestamps
	if err := json.Unmarshal([]byte(`{
  "enable_logging": true,
  "output_format": "value",
  "inputs": [
    {
      "text": "value",
      "voice_id": "value"
    }
  ],
  "model_id": "value",
  "language_code": "value",
  "settings": {
    "stability": 0.5,
    "similarity": 0.5,
    "clone_strength": 0.5
  },
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "seed": 123,
  "apply_text_normalization": "auto",
  "use_pvc_as_ivc": true,
  "previous_request_ids": [
    "value"
  ],
  "next_request_ids": [
    "value"
  ],
  "previous_text": "value",
  "future_text": "value",
  "audio_effects": {
    "filter_preset_id": "value",
    "distance": 0.5,
    "environment_id": "value",
    "background_noise_id": "value",
    "send_level": 0.5,
    "pan": 0.5
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	outputFormatValue := string("value")
	request.OutputFormat = &outputFormatValue
	enableLoggingValue := bool(true)
	request.EnableLogging = &enableLoggingValue
	_, err := client.FullWithTimestamps(context.Background(), &request)
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

type wireRequestTexttodialogue struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireTexttodialogue records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireTexttodialogue(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestTexttodialogue) {
	t.Helper()
	recorded := make(chan wireRequestTexttodialogue, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestTexttodialogue{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireTexttodialogue(t *testing.T, recorded <-chan wireRequestTexttodialogue) wireRequestTexttodialogue {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestTexttodialogue{}
	}
}

// wireJSONContainsTexttodialogue reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsTexttodialogue(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualTexttodialogue(got, want)
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

// setWireValueTexttodialogue sets a named string, int or float field, or the value a pointer field points to.
func setWireValueTexttodialogue(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualTexttodialogue(got, want []byte) bool {
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
