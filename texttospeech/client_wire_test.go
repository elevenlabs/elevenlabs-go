package texttospeech

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

func TestConvertWithTimestampsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireTexttospeech(t, 200, `{
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
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodyTextToSpeechFullWithTimestamps
	if err := json.Unmarshal([]byte(`{
  "output_format": "alaw_8000",
  "optimize_streaming_latency": 123,
  "enable_logging": true,
  "voice_id": "value",
  "text": "value",
  "model_id": "value",
  "language_code": "value",
  "voice_settings": {
    "stability": 0.5,
    "use_speaker_boost": true,
    "similarity_boost": 0.5,
    "style": 9.99,
    "speed": 9.99,
    "clone_strength": 0.5
  },
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "seed": 123,
  "previous_text": "value",
  "next_text": "value",
  "previous_request_ids": [
    "value"
  ],
  "next_request_ids": [
    "value"
  ],
  "use_pvc_as_ivc": true,
  "apply_text_normalization": "auto",
  "apply_language_text_normalization": true,
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
	enableLoggingValue := bool(true)
	request.EnableLogging = &enableLoggingValue
	optimizeStreamingLatencyValue := int(123)
	request.OptimizeStreamingLatency = &optimizeStreamingLatencyValue
	result, err := client.ConvertWithTimestamps(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualTexttospeech(gotResponse, []byte(`{
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
  }
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
  }
}`)
	}

	req := receiveWireTexttospeech(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/text-to-speech/value/with-timestamps" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/text-to-speech/value/with-timestamps")
	}
	if got := req.Query.Get("enable_logging"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "enable_logging", got, "true")
	}
	if got := req.Query.Get("optimize_streaming_latency"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "optimize_streaming_latency", got, "123")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsTexttospeech(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestConvertWithTimestampsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireTexttospeech(t, 422, `{
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

	var request types.BodyTextToSpeechFullWithTimestamps
	if err := json.Unmarshal([]byte(`{
  "output_format": "alaw_8000",
  "optimize_streaming_latency": 123,
  "enable_logging": true,
  "voice_id": "value",
  "text": "value",
  "model_id": "value",
  "language_code": "value",
  "voice_settings": {
    "stability": 0.5,
    "use_speaker_boost": true,
    "similarity_boost": 0.5,
    "style": 9.99,
    "speed": 9.99,
    "clone_strength": 0.5
  },
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ],
  "seed": 123,
  "previous_text": "value",
  "next_text": "value",
  "previous_request_ids": [
    "value"
  ],
  "next_request_ids": [
    "value"
  ],
  "use_pvc_as_ivc": true,
  "apply_text_normalization": "auto",
  "apply_language_text_normalization": true,
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
	enableLoggingValue := bool(true)
	request.EnableLogging = &enableLoggingValue
	optimizeStreamingLatencyValue := int(123)
	request.OptimizeStreamingLatency = &optimizeStreamingLatencyValue
	_, err := client.ConvertWithTimestamps(context.Background(), string("value"), &request)
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

type wireRequestTexttospeech struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireTexttospeech records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireTexttospeech(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestTexttospeech) {
	t.Helper()
	recorded := make(chan wireRequestTexttospeech, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestTexttospeech{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireTexttospeech(t *testing.T, recorded <-chan wireRequestTexttospeech) wireRequestTexttospeech {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestTexttospeech{}
	}
}

// wireJSONContainsTexttospeech reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsTexttospeech(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualTexttospeech(got, want)
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

// setWireValueTexttospeech sets a named string, int or float field, or the value a pointer field points to.
func setWireValueTexttospeech(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualTexttospeech(got, want []byte) bool {
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
