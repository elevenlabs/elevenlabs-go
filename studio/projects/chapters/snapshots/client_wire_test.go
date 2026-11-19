package snapshots

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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireStudioProjectsChaptersSnapshots(t, 200, `{
  "snapshots": [
    {
      "chapter_snapshot_id": "value",
      "project_id": "value",
      "chapter_id": "value",
      "created_at_unix": 123,
      "name": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.List(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsChaptersSnapshots(gotResponse, []byte(`{
  "snapshots": [
    {
      "chapter_snapshot_id": "value",
      "project_id": "value",
      "chapter_id": "value",
      "created_at_unix": 123,
      "name": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "snapshots": [
    {
      "chapter_snapshot_id": "value",
      "project_id": "value",
      "chapter_id": "value",
      "created_at_unix": 123,
      "name": "value"
    }
  ]
}`)
	}

	req := receiveWireStudioProjectsChaptersSnapshots(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/studio/projects/value/chapters/value2/snapshots" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/chapters/value2/snapshots")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsChaptersSnapshots(t, 422, `{
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

	_, err := client.List(context.Background(), string("value"), string("value2"))
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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireStudioProjectsChaptersSnapshots(t, 200, `{
  "chapter_snapshot_id": "value",
  "project_id": "value",
  "chapter_id": "value",
  "created_at_unix": 123,
  "name": "value",
  "character_alignments": [
    {
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
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"), string("value2"), string("value3"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsChaptersSnapshots(gotResponse, []byte(`{
  "chapter_snapshot_id": "value",
  "project_id": "value",
  "chapter_id": "value",
  "created_at_unix": 123,
  "name": "value",
  "character_alignments": [
    {
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
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "chapter_snapshot_id": "value",
  "project_id": "value",
  "chapter_id": "value",
  "created_at_unix": 123,
  "name": "value",
  "character_alignments": [
    {
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
  ]
}`)
	}

	req := receiveWireStudioProjectsChaptersSnapshots(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/studio/projects/value/chapters/value2/snapshots/value3" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/chapters/value2/snapshots/value3")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsChaptersSnapshots(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"), string("value2"), string("value3"))
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

type wireRequestStudioProjectsChaptersSnapshots struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireStudioProjectsChaptersSnapshots records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireStudioProjectsChaptersSnapshots(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestStudioProjectsChaptersSnapshots) {
	t.Helper()
	recorded := make(chan wireRequestStudioProjectsChaptersSnapshots, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestStudioProjectsChaptersSnapshots{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireStudioProjectsChaptersSnapshots(t *testing.T, recorded <-chan wireRequestStudioProjectsChaptersSnapshots) wireRequestStudioProjectsChaptersSnapshots {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestStudioProjectsChaptersSnapshots{}
	}
}

// wireJSONContainsStudioProjectsChaptersSnapshots reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsStudioProjectsChaptersSnapshots(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualStudioProjectsChaptersSnapshots(got, want)
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

// setWireValueStudioProjectsChaptersSnapshots sets a named string, int or float field, or the value a pointer field points to.
func setWireValueStudioProjectsChaptersSnapshots(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualStudioProjectsChaptersSnapshots(got, want []byte) bool {
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
