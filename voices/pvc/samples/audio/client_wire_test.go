package audio

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
	samples "github.com/elevenlabs/elevenlabs-go/voices/pvc/samples"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesPvcSamplesAudio(t, 200, `{
  "audio_base_64": "value",
  "voice_id": "value",
  "sample_id": "value",
  "media_type": "value",
  "duration_secs": 9.99
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	removeBackgroundNoiseValue := bool(true)
	request := &samples.GetAudioRequest{
		RemoveBackgroundNoise: &removeBackgroundNoiseValue,
	}
	result, err := client.Get(context.Background(), string("value"), string("value2"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesPvcSamplesAudio(gotResponse, []byte(`{
  "audio_base_64": "value",
  "voice_id": "value",
  "sample_id": "value",
  "media_type": "value",
  "duration_secs": 9.99
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "audio_base_64": "value",
  "voice_id": "value",
  "sample_id": "value",
  "media_type": "value",
  "duration_secs": 9.99
}`)
	}

	req := receiveWireVoicesPvcSamplesAudio(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/voices/pvc/value/samples/value2/audio" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/pvc/value/samples/value2/audio")
	}
	if got := req.Query.Get("remove_background_noise"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "remove_background_noise", got, "true")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesPvcSamplesAudio(t, 422, `{
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

	removeBackgroundNoiseValue := bool(true)
	request := &samples.GetAudioRequest{
		RemoveBackgroundNoise: &removeBackgroundNoiseValue,
	}
	_, err := client.Get(context.Background(), string("value"), string("value2"), request)
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

type wireRequestVoicesPvcSamplesAudio struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireVoicesPvcSamplesAudio records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireVoicesPvcSamplesAudio(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestVoicesPvcSamplesAudio) {
	t.Helper()
	recorded := make(chan wireRequestVoicesPvcSamplesAudio, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestVoicesPvcSamplesAudio{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireVoicesPvcSamplesAudio(t *testing.T, recorded <-chan wireRequestVoicesPvcSamplesAudio) wireRequestVoicesPvcSamplesAudio {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestVoicesPvcSamplesAudio{}
	}
}

// wireJSONContainsVoicesPvcSamplesAudio reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsVoicesPvcSamplesAudio(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualVoicesPvcSamplesAudio(got, want)
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

// setWireValueVoicesPvcSamplesAudio sets a named string, int or float field, or the value a pointer field points to.
func setWireValueVoicesPvcSamplesAudio(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualVoicesPvcSamplesAudio(got, want []byte) bool {
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
