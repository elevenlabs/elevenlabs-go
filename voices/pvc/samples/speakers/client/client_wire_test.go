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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesPvcSamplesSpeakersClient(t, 200, `{
  "voice_id": "value",
  "sample_id": "value",
  "status": "not_started",
  "speakers": {
    "key": {
      "speaker_id": "value",
      "duration_secs": 9.99,
      "utterances": [
        {
          "start": 9.99,
          "end": 9.99
        }
      ]
    }
  },
  "selected_speaker_ids": [
    "value"
  ]
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
	if !wireJSONBytesEqualVoicesPvcSamplesSpeakersClient(gotResponse, []byte(`{
  "voice_id": "value",
  "sample_id": "value",
  "status": "not_started",
  "speakers": {
    "key": {
      "speaker_id": "value",
      "duration_secs": 9.99,
      "utterances": [
        {
          "start": 9.99,
          "end": 9.99
        }
      ]
    }
  },
  "selected_speaker_ids": [
    "value"
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "voice_id": "value",
  "sample_id": "value",
  "status": "not_started",
  "speakers": {
    "key": {
      "speaker_id": "value",
      "duration_secs": 9.99,
      "utterances": [
        {
          "start": 9.99,
          "end": 9.99
        }
      ]
    }
  },
  "selected_speaker_ids": [
    "value"
  ]
}`)
	}

	req := receiveWireVoicesPvcSamplesSpeakersClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/voices/pvc/value/samples/value2/speakers" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/pvc/value/samples/value2/speakers")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesPvcSamplesSpeakersClient(t, 422, `{
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

func TestSeparateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesPvcSamplesSpeakersClient(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Separate(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesPvcSamplesSpeakersClient(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireVoicesPvcSamplesSpeakersClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/voices/pvc/value/samples/value2/separate-speakers" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/pvc/value/samples/value2/separate-speakers")
	}
}

func TestSeparateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesPvcSamplesSpeakersClient(t, 422, `{
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

	_, err := client.Separate(context.Background(), string("value"), string("value2"))
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

type wireRequestVoicesPvcSamplesSpeakersClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireVoicesPvcSamplesSpeakersClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireVoicesPvcSamplesSpeakersClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestVoicesPvcSamplesSpeakersClient) {
	t.Helper()
	recorded := make(chan wireRequestVoicesPvcSamplesSpeakersClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestVoicesPvcSamplesSpeakersClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireVoicesPvcSamplesSpeakersClient(t *testing.T, recorded <-chan wireRequestVoicesPvcSamplesSpeakersClient) wireRequestVoicesPvcSamplesSpeakersClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestVoicesPvcSamplesSpeakersClient{}
	}
}

// wireJSONContainsVoicesPvcSamplesSpeakersClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsVoicesPvcSamplesSpeakersClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualVoicesPvcSamplesSpeakersClient(got, want)
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

// setWireValueVoicesPvcSamplesSpeakersClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueVoicesPvcSamplesSpeakersClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualVoicesPvcSamplesSpeakersClient(got, want []byte) bool {
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
