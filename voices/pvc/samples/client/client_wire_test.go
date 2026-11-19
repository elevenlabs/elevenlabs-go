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
	pvc "github.com/elevenlabs/elevenlabs-go/voices/pvc"
)

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesPvcSamplesClient(t, 200, `{
  "voice_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request pvc.BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost
	if err := json.Unmarshal([]byte(`{
  "sample_id": "value",
  "voice_id": "value",
  "remove_background_noise": true,
  "selected_speaker_ids": [
    "value"
  ],
  "trim_start_time": 123,
  "trim_end_time": 123,
  "file_name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesPvcSamplesClient(gotResponse, []byte(`{
  "voice_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "voice_id": "value"
}`)
	}

	req := receiveWireVoicesPvcSamplesClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/voices/pvc/value/samples/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/pvc/value/samples/value2")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsVoicesPvcSamplesClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesPvcSamplesClient(t, 422, `{
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

	var request pvc.BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost
	if err := json.Unmarshal([]byte(`{
  "sample_id": "value",
  "voice_id": "value",
  "remove_background_noise": true,
  "selected_speaker_ids": [
    "value"
  ],
  "trim_start_time": 123,
  "trim_end_time": 123,
  "file_name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), string("value2"), &request)
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

	server, recorded := serveWireVoicesPvcSamplesClient(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesPvcSamplesClient(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireVoicesPvcSamplesClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/voices/pvc/value/samples/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/pvc/value/samples/value2")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesPvcSamplesClient(t, 422, `{
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

	_, err := client.Delete(context.Background(), string("value"), string("value2"))
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

type wireRequestVoicesPvcSamplesClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireVoicesPvcSamplesClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireVoicesPvcSamplesClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestVoicesPvcSamplesClient) {
	t.Helper()
	recorded := make(chan wireRequestVoicesPvcSamplesClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestVoicesPvcSamplesClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireVoicesPvcSamplesClient(t *testing.T, recorded <-chan wireRequestVoicesPvcSamplesClient) wireRequestVoicesPvcSamplesClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestVoicesPvcSamplesClient{}
	}
}

// wireJSONContainsVoicesPvcSamplesClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsVoicesPvcSamplesClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualVoicesPvcSamplesClient(got, want)
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

// setWireValueVoicesPvcSamplesClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueVoicesPvcSamplesClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualVoicesPvcSamplesClient(got, want []byte) bool {
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
