package accents

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
	voices "github.com/elevenlabs/elevenlabs-go/voices"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireVoicesAccents(t, 200, `{
  "accents": [
    {
      "accent": "value",
      "language": "value",
      "code": "value",
      "name": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	languageValue := string("value")
	modelIdValue := string("value")
	request := &voices.GetAccentsRequest{
		Language: &languageValue,
		ModelId:  &modelIdValue,
	}
	result, err := client.Get(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualVoicesAccents(gotResponse, []byte(`{
  "accents": [
    {
      "accent": "value",
      "language": "value",
      "code": "value",
      "name": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "accents": [
    {
      "accent": "value",
      "language": "value",
      "code": "value",
      "name": "value"
    }
  ]
}`)
	}

	req := receiveWireVoicesAccents(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/voices/accents" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/voices/accents")
	}
	if got := req.Query.Get("language"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "language", got, "value")
	}
	if got := req.Query.Get("model_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "model_id", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireVoicesAccents(t, 422, `{
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

	languageValue := string("value")
	modelIdValue := string("value")
	request := &voices.GetAccentsRequest{
		Language: &languageValue,
		ModelId:  &modelIdValue,
	}
	_, err := client.Get(context.Background(), request)
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

type wireRequestVoicesAccents struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireVoicesAccents records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireVoicesAccents(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestVoicesAccents) {
	t.Helper()
	recorded := make(chan wireRequestVoicesAccents, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestVoicesAccents{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireVoicesAccents(t *testing.T, recorded <-chan wireRequestVoicesAccents) wireRequestVoicesAccents {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestVoicesAccents{}
	}
}

// wireJSONContainsVoicesAccents reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsVoicesAccents(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualVoicesAccents(got, want)
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

// setWireValueVoicesAccents sets a named string, int or float field, or the value a pointer field points to.
func setWireValueVoicesAccents(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualVoicesAccents(got, want []byte) bool {
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
