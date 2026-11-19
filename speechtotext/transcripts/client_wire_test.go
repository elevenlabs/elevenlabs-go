package transcripts

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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireSpeechtotextTranscripts(t, 200, `{
  "language_code": "value",
  "language_probability": 9.99,
  "text": "value",
  "words": [
    {
      "text": "value",
      "start": 9.99,
      "end": 9.99,
      "type": "word",
      "speaker_id": "value",
      "logprob": 9.99,
      "characters": [
        {
          "text": "value",
          "start": 9.99,
          "end": 9.99
        }
      ],
      "channel_index": 123
    }
  ],
  "channel_index": 123,
  "additional_formats": [
    {
      "requested_format": "value",
      "file_extension": "value",
      "content_type": "value",
      "is_base64_encoded": true,
      "content": "value"
    }
  ],
  "transcription_id": "value",
  "entities": [
    {
      "text": "value",
      "entity_type": "value",
      "start_char": 123,
      "end_char": 123
    }
  ],
  "audio_duration_secs": 9.99,
  "edited_transcript": {
    "kind": "transcript",
    "text": "value",
    "edited_text": "value"
  }
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
	if !wireJSONBytesEqualSpeechtotextTranscripts(gotResponse, []byte(`{
  "language_code": "value",
  "language_probability": 9.99,
  "text": "value",
  "words": [
    {
      "text": "value",
      "start": 9.99,
      "end": 9.99,
      "type": "word",
      "speaker_id": "value",
      "logprob": 9.99,
      "characters": [
        {
          "text": "value",
          "start": 9.99,
          "end": 9.99
        }
      ],
      "channel_index": 123
    }
  ],
  "channel_index": 123,
  "additional_formats": [
    {
      "requested_format": "value",
      "file_extension": "value",
      "content_type": "value",
      "is_base64_encoded": true,
      "content": "value"
    }
  ],
  "transcription_id": "value",
  "entities": [
    {
      "text": "value",
      "entity_type": "value",
      "start_char": 123,
      "end_char": 123
    }
  ],
  "audio_duration_secs": 9.99,
  "edited_transcript": {
    "kind": "transcript",
    "text": "value",
    "edited_text": "value"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "language_code": "value",
  "language_probability": 9.99,
  "text": "value",
  "words": [
    {
      "text": "value",
      "start": 9.99,
      "end": 9.99,
      "type": "word",
      "speaker_id": "value",
      "logprob": 9.99,
      "characters": [
        {
          "text": "value",
          "start": 9.99,
          "end": 9.99
        }
      ],
      "channel_index": 123
    }
  ],
  "channel_index": 123,
  "additional_formats": [
    {
      "requested_format": "value",
      "file_extension": "value",
      "content_type": "value",
      "is_base64_encoded": true,
      "content": "value"
    }
  ],
  "transcription_id": "value",
  "entities": [
    {
      "text": "value",
      "entity_type": "value",
      "start_char": 123,
      "end_char": 123
    }
  ],
  "audio_duration_secs": 9.99,
  "edited_transcript": {
    "kind": "transcript",
    "text": "value",
    "edited_text": "value"
  }
}`)
	}

	req := receiveWireSpeechtotextTranscripts(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/speech-to-text/transcripts/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/speech-to-text/transcripts/value")
	}
}

func TestGetWireError401(t *testing.T) {
	t.Parallel()

	server, _ := serveWireSpeechtotextTranscripts(t, 401, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Get(context.Background(), string("value"))
	if err == nil {
		t.Fatal("expected an error for a 401 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 401 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 401)
	}
}

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireSpeechtotextTranscripts(t, 200, `{}`)

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
	if !wireJSONBytesEqualSpeechtotextTranscripts(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireSpeechtotextTranscripts(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/speech-to-text/transcripts/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/speech-to-text/transcripts/value")
	}
}

func TestDeleteWireError401(t *testing.T) {
	t.Parallel()

	server, _ := serveWireSpeechtotextTranscripts(t, 401, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	_, err := client.Delete(context.Background(), string("value"))
	if err == nil {
		t.Fatal("expected an error for a 401 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 401 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 401)
	}
}

type wireRequestSpeechtotextTranscripts struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireSpeechtotextTranscripts records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireSpeechtotextTranscripts(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestSpeechtotextTranscripts) {
	t.Helper()
	recorded := make(chan wireRequestSpeechtotextTranscripts, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestSpeechtotextTranscripts{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireSpeechtotextTranscripts(t *testing.T, recorded <-chan wireRequestSpeechtotextTranscripts) wireRequestSpeechtotextTranscripts {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestSpeechtotextTranscripts{}
	}
}

// wireJSONContainsSpeechtotextTranscripts reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsSpeechtotextTranscripts(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualSpeechtotextTranscripts(got, want)
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

// setWireValueSpeechtotextTranscripts sets a named string, int or float field, or the value a pointer field points to.
func setWireValueSpeechtotextTranscripts(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualSpeechtotextTranscripts(got, want []byte) bool {
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
