package audionative

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

func TestGetSettingsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAudionative(t, 200, `{
  "enabled": true,
  "snapshot_id": "value",
  "settings": {
    "title": "value",
    "image": "value",
    "author": "value",
    "small": true,
    "text_color": "value",
    "background_color": "value",
    "sessionization": 123,
    "audio_path": "value",
    "audio_url": "value",
    "status": "processing"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetSettings(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAudionative(gotResponse, []byte(`{
  "enabled": true,
  "snapshot_id": "value",
  "settings": {
    "title": "value",
    "image": "value",
    "author": "value",
    "small": true,
    "text_color": "value",
    "background_color": "value",
    "sessionization": 123,
    "audio_path": "value",
    "audio_url": "value",
    "status": "processing"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "enabled": true,
  "snapshot_id": "value",
  "settings": {
    "title": "value",
    "image": "value",
    "author": "value",
    "small": true,
    "text_color": "value",
    "background_color": "value",
    "sessionization": 123,
    "audio_path": "value",
    "audio_url": "value",
    "status": "processing"
  }
}`)
	}

	req := receiveWireAudionative(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/audio-native/value/settings" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/audio-native/value/settings")
	}
}

func TestGetSettingsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAudionative(t, 422, `{
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

	_, err := client.GetSettings(context.Background(), string("value"))
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

func TestUpdateContentFromUrlWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAudionative(t, 200, `{
  "project_id": "value",
  "converting": true,
  "publishing": true,
  "html_snippet": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost
	if err := json.Unmarshal([]byte(`{
  "url": "value",
  "author": "value",
  "title": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.UpdateContentFromUrl(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAudionative(gotResponse, []byte(`{
  "project_id": "value",
  "converting": true,
  "publishing": true,
  "html_snippet": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "project_id": "value",
  "converting": true,
  "publishing": true,
  "html_snippet": "value"
}`)
	}

	req := receiveWireAudionative(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/audio-native/content" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/audio-native/content")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAudionative(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateContentFromUrlWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAudionative(t, 422, `{
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

	var request types.BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost
	if err := json.Unmarshal([]byte(`{
  "url": "value",
  "author": "value",
  "title": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.UpdateContentFromUrl(context.Background(), &request)
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

type wireRequestAudionative struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAudionative records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAudionative(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAudionative) {
	t.Helper()
	recorded := make(chan wireRequestAudionative, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAudionative{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAudionative(t *testing.T, recorded <-chan wireRequestAudionative) wireRequestAudionative {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAudionative{}
	}
}

// wireJSONContainsAudionative reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAudionative(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAudionative(got, want)
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

// setWireValueAudionative sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAudionative(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAudionative(got, want []byte) bool {
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
