package settings

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

func TestGetDefaultWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesSettings(t, 200, `{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetDefault(context.Background())
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesSettings(gotResponse, []byte(`{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`)
	}

	req := receiveWireVoicesSettings(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/voices/settings/default" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/settings/default")
	}
}

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesSettings(t, 200, `{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
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
	if !wireJSONBytesEqualVoicesSettings(gotResponse, []byte(`{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`)
	}

	req := receiveWireVoicesSettings(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/voices/value/settings" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/value/settings")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesSettings(t, 422, `{
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesSettings(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.VoiceSettingsResponseModel
	if err := json.Unmarshal([]byte(`{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesSettings(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireVoicesSettings(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/voices/value/settings/edit" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/value/settings/edit")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsVoicesSettings(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesSettings(t, 422, `{
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

	var request types.VoiceSettingsResponseModel
	if err := json.Unmarshal([]byte(`{
  "stability": 0.5,
  "use_speaker_boost": true,
  "similarity_boost": 0.5,
  "style": 9.99,
  "speed": 9.99,
  "clone_strength": 0.5
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), &request)
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

type wireRequestVoicesSettings struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireVoicesSettings records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireVoicesSettings(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestVoicesSettings) {
	t.Helper()
	recorded := make(chan wireRequestVoicesSettings, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestVoicesSettings{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireVoicesSettings(t *testing.T, recorded <-chan wireRequestVoicesSettings) wireRequestVoicesSettings {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestVoicesSettings{}
	}
}

// wireJSONContainsVoicesSettings reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsVoicesSettings(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualVoicesSettings(got, want)
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

// setWireValueVoicesSettings sets a named string, int or float field, or the value a pointer field points to.
func setWireValueVoicesSettings(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualVoicesSettings(got, want []byte) bool {
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
