package transcript

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
	language "github.com/elevenlabs/elevenlabs-go/dubbing/project/language"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectLanguageTranscript(t, 200, `{
  "source_language": "value",
  "target_language": "value",
  "segments": [
    {
      "id": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "source_text": "value",
      "translation": "value"
    }
  ],
  "revision": 123
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
	if !wireJSONBytesEqualDubbingProjectLanguageTranscript(gotResponse, []byte(`{
  "source_language": "value",
  "target_language": "value",
  "segments": [
    {
      "id": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "source_text": "value",
      "translation": "value"
    }
  ],
  "revision": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "source_language": "value",
  "target_language": "value",
  "segments": [
    {
      "id": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "source_text": "value",
      "translation": "value"
    }
  ],
  "revision": 123
}`)
	}

	req := receiveWireDubbingProjectLanguageTranscript(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/dubbing/project/value/language/value2/transcript" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/language/value2/transcript")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectLanguageTranscript(t, 422, `{
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

func TestUpdateSegmentWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectLanguageTranscript(t, 200, `{
  "segment": {
    "id": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "source_text": "value",
    "translation": "value"
  },
  "revision": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.DubbingTargetSegmentUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "segment_id": "value",
  "language_id": "value",
  "project_id": "value",
  "translation": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.UpdateSegment(context.Background(), string("value"), string("value2"), string("value3"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectLanguageTranscript(gotResponse, []byte(`{
  "segment": {
    "id": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "source_text": "value",
    "translation": "value"
  },
  "revision": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "segment": {
    "id": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "source_text": "value",
    "translation": "value"
  },
  "revision": 123
}`)
	}

	req := receiveWireDubbingProjectLanguageTranscript(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/dubbing/project/value/language/value2/transcript/segment/value3" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/language/value2/transcript/segment/value3")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsDubbingProjectLanguageTranscript(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateSegmentWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectLanguageTranscript(t, 422, `{
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

	var request types.DubbingTargetSegmentUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "segment_id": "value",
  "language_id": "value",
  "project_id": "value",
  "translation": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.UpdateSegment(context.Background(), string("value"), string("value2"), string("value3"), &request)
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

func TestUpdateSegmentsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectLanguageTranscript(t, 200, `{
  "segments": [
    {
      "id": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "source_text": "value",
      "translation": "value"
    }
  ],
  "revision": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request language.DubbingBulkTargetSegmentUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "language_id": "value",
  "project_id": "value",
  "segments": {
    "key": {
      "segment_id": "value",
      "language_id": "value",
      "project_id": "value",
      "translation": "value"
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.UpdateSegments(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectLanguageTranscript(gotResponse, []byte(`{
  "segments": [
    {
      "id": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "source_text": "value",
      "translation": "value"
    }
  ],
  "revision": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "segments": [
    {
      "id": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "source_text": "value",
      "translation": "value"
    }
  ],
  "revision": 123
}`)
	}

	req := receiveWireDubbingProjectLanguageTranscript(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/dubbing/project/value/language/value2/transcript/segments" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/language/value2/transcript/segments")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsDubbingProjectLanguageTranscript(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateSegmentsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectLanguageTranscript(t, 422, `{
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

	var request language.DubbingBulkTargetSegmentUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "language_id": "value",
  "project_id": "value",
  "segments": {
    "key": {
      "segment_id": "value",
      "language_id": "value",
      "project_id": "value",
      "translation": "value"
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.UpdateSegments(context.Background(), string("value"), string("value2"), &request)
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

func TestRegenerateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectLanguageTranscript(t, 202, `{
  "regenerated_segment_ids": [
    "value"
  ],
  "regenerated_seconds": 9.99,
  "charged_seconds": 9.99,
  "free_regeneration_seconds_remaining": 9.99
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Regenerate(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectLanguageTranscript(gotResponse, []byte(`{
  "regenerated_segment_ids": [
    "value"
  ],
  "regenerated_seconds": 9.99,
  "charged_seconds": 9.99,
  "free_regeneration_seconds_remaining": 9.99
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "regenerated_segment_ids": [
    "value"
  ],
  "regenerated_seconds": 9.99,
  "charged_seconds": 9.99,
  "free_regeneration_seconds_remaining": 9.99
}`)
	}

	req := receiveWireDubbingProjectLanguageTranscript(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/dubbing/project/value/language/value2/transcript/regenerate" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/language/value2/transcript/regenerate")
	}
}

func TestRegenerateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectLanguageTranscript(t, 422, `{
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

	_, err := client.Regenerate(context.Background(), string("value"), string("value2"))
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

type wireRequestDubbingProjectLanguageTranscript struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireDubbingProjectLanguageTranscript records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireDubbingProjectLanguageTranscript(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestDubbingProjectLanguageTranscript) {
	t.Helper()
	recorded := make(chan wireRequestDubbingProjectLanguageTranscript, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestDubbingProjectLanguageTranscript{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireDubbingProjectLanguageTranscript(t *testing.T, recorded <-chan wireRequestDubbingProjectLanguageTranscript) wireRequestDubbingProjectLanguageTranscript {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestDubbingProjectLanguageTranscript{}
	}
}

// wireJSONContainsDubbingProjectLanguageTranscript reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsDubbingProjectLanguageTranscript(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualDubbingProjectLanguageTranscript(got, want)
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

// setWireValueDubbingProjectLanguageTranscript sets a named string, int or float field, or the value a pointer field points to.
func setWireValueDubbingProjectLanguageTranscript(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualDubbingProjectLanguageTranscript(got, want []byte) bool {
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
