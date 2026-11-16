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
	project "github.com/elevenlabs/elevenlabs-go/dubbing/project"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectTranscript(t, 200, `{
  "language": "value",
  "segments": [
    {
      "id": "value",
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "external_id": "value"
    }
  ],
  "revision": 123
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
	if !wireJSONBytesEqualDubbingProjectTranscript(gotResponse, []byte(`{
  "language": "value",
  "segments": [
    {
      "id": "value",
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "external_id": "value"
    }
  ],
  "revision": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "language": "value",
  "segments": [
    {
      "id": "value",
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "external_id": "value"
    }
  ],
  "revision": 123
}`)
	}

	req := receiveWireDubbingProjectTranscript(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/dubbing/project/value/transcript" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/transcript")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectTranscript(t, 422, `{
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

func TestUpdateSegmentWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectTranscript(t, 200, `{
  "segment": {
    "id": "value",
    "text": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "external_id": "value"
  },
  "revision": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.DubbingSegmentUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "segment_id": "value",
  "project_id": "value",
  "text": "value",
  "speaker_id": "value",
  "start_s": 9.99,
  "end_s": 9.99
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.UpdateSegment(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectTranscript(gotResponse, []byte(`{
  "segment": {
    "id": "value",
    "text": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "external_id": "value"
  },
  "revision": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "segment": {
    "id": "value",
    "text": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "external_id": "value"
  },
  "revision": 123
}`)
	}

	req := receiveWireDubbingProjectTranscript(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/dubbing/project/value/transcript/segment/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/transcript/segment/value2")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsDubbingProjectTranscript(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateSegmentWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectTranscript(t, 422, `{
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

	var request types.DubbingSegmentUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "segment_id": "value",
  "project_id": "value",
  "text": "value",
  "speaker_id": "value",
  "start_s": 9.99,
  "end_s": 9.99
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.UpdateSegment(context.Background(), string("value"), string("value2"), &request)
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

func TestDeleteSegmentWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectTranscript(t, 200, `{
  "revision": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.DeleteSegment(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectTranscript(gotResponse, []byte(`{
  "revision": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "revision": 123
}`)
	}

	req := receiveWireDubbingProjectTranscript(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/dubbing/project/value/transcript/segment/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/transcript/segment/value2")
	}
}

func TestDeleteSegmentWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectTranscript(t, 422, `{
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

	_, err := client.DeleteSegment(context.Background(), string("value"), string("value2"))
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

	server, recorded := serveWireDubbingProjectTranscript(t, 200, `{
  "segments": [
    {
      "id": "value",
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "external_id": "value"
    }
  ],
  "revision": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request project.DubbingBulkSegmentUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "segments": {
    "key": {
      "segment_id": "value",
      "project_id": "value",
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.UpdateSegments(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectTranscript(gotResponse, []byte(`{
  "segments": [
    {
      "id": "value",
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "external_id": "value"
    }
  ],
  "revision": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "segments": [
    {
      "id": "value",
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "external_id": "value"
    }
  ],
  "revision": 123
}`)
	}

	req := receiveWireDubbingProjectTranscript(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/dubbing/project/value/transcript/segments" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/transcript/segments")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsDubbingProjectTranscript(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateSegmentsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectTranscript(t, 422, `{
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

	var request project.DubbingBulkSegmentUpdateRequest
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "segments": {
    "key": {
      "segment_id": "value",
      "project_id": "value",
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99
    }
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.UpdateSegments(context.Background(), string("value"), &request)
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

func TestCreateSegmentWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireDubbingProjectTranscript(t, 201, `{
  "segment": {
    "id": "value",
    "text": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "external_id": "value"
  },
  "revision": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request project.DubbingSegmentCreateRequest
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "text": "value",
  "speaker_id": "value",
  "start_s": 9.99,
  "end_s": 9.99
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CreateSegment(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualDubbingProjectTranscript(gotResponse, []byte(`{
  "segment": {
    "id": "value",
    "text": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "external_id": "value"
  },
  "revision": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "segment": {
    "id": "value",
    "text": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "external_id": "value"
  },
  "revision": 123
}`)
	}

	req := receiveWireDubbingProjectTranscript(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/dubbing/project/value/transcript/segment" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/dubbing/project/value/transcript/segment")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsDubbingProjectTranscript(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateSegmentWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireDubbingProjectTranscript(t, 422, `{
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

	var request project.DubbingSegmentCreateRequest
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "text": "value",
  "speaker_id": "value",
  "start_s": 9.99,
  "end_s": 9.99
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CreateSegment(context.Background(), string("value"), &request)
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

type wireRequestDubbingProjectTranscript struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireDubbingProjectTranscript records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireDubbingProjectTranscript(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestDubbingProjectTranscript) {
	t.Helper()
	recorded := make(chan wireRequestDubbingProjectTranscript, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestDubbingProjectTranscript{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireDubbingProjectTranscript(t *testing.T, recorded <-chan wireRequestDubbingProjectTranscript) wireRequestDubbingProjectTranscript {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestDubbingProjectTranscript{}
	}
}

// wireJSONContainsDubbingProjectTranscript reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsDubbingProjectTranscript(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualDubbingProjectTranscript(got, want)
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

// setWireValueDubbingProjectTranscript sets a named string, int or float field, or the value a pointer field points to.
func setWireValueDubbingProjectTranscript(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualDubbingProjectTranscript(got, want []byte) bool {
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
