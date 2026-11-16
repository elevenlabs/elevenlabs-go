package runs

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
	templates "github.com/elevenlabs/elevenlabs-go/flows/templates"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireFlowsTemplatesRuns(t, 200, `{
  "runs": [
    {
      "id": "value",
      "template_id": "value",
      "version_id": "value",
      "status": "pending",
      "outputs": {
        "key": {
          "id": "value",
          "status": "pending",
          "failure_reason": "timeout",
          "error_message": "value",
          "content_url": "value",
          "content_mime_type": "value",
          "type": "image"
        }
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(100)
	versionIdValue := string("value")
	request := &templates.ListRunsRequest{
		Cursor:    &cursorValue,
		PageSize:  &pageSizeValue,
		VersionId: &versionIdValue,
	}
	result, err := client.List(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualFlowsTemplatesRuns(gotResponse, []byte(`{
  "runs": [
    {
      "id": "value",
      "template_id": "value",
      "version_id": "value",
      "status": "pending",
      "outputs": {
        "key": {
          "id": "value",
          "status": "pending",
          "failure_reason": "timeout",
          "error_message": "value",
          "content_url": "value",
          "content_mime_type": "value",
          "type": "image"
        }
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "runs": [
    {
      "id": "value",
      "template_id": "value",
      "version_id": "value",
      "status": "pending",
      "outputs": {
        "key": {
          "id": "value",
          "status": "pending",
          "failure_reason": "timeout",
          "error_message": "value",
          "content_url": "value",
          "content_mime_type": "value",
          "type": "image"
        }
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireFlowsTemplatesRuns(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/flows/templates/value/runs" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/flows/templates/value/runs")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("version_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "version_id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireFlowsTemplatesRuns(t, 422, `{
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

	cursorValue := string("value")
	pageSizeValue := int(100)
	versionIdValue := string("value")
	request := &templates.ListRunsRequest{
		Cursor:    &cursorValue,
		PageSize:  &pageSizeValue,
		VersionId: &versionIdValue,
	}
	_, err := client.List(context.Background(), string("value"), request)
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireFlowsTemplatesRuns(t, 200, `{
  "id": "value",
  "template_id": "value",
  "version_id": "value",
  "status": "pending",
  "outputs": {
    "key": {
      "id": "value",
      "status": "pending",
      "failure_reason": "timeout",
      "error_message": "value",
      "content_url": "value",
      "content_mime_type": "value",
      "type": "image"
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request templates.TemplateRunCreateRequest
	if err := json.Unmarshal([]byte(`{
  "template_id": "value",
  "inputs": {
    "key": "value"
  },
  "version_id": "value",
  "webhook": {
    "type": "all"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Create(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualFlowsTemplatesRuns(gotResponse, []byte(`{
  "id": "value",
  "template_id": "value",
  "version_id": "value",
  "status": "pending",
  "outputs": {
    "key": {
      "id": "value",
      "status": "pending",
      "failure_reason": "timeout",
      "error_message": "value",
      "content_url": "value",
      "content_mime_type": "value",
      "type": "image"
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "template_id": "value",
  "version_id": "value",
  "status": "pending",
  "outputs": {
    "key": {
      "id": "value",
      "status": "pending",
      "failure_reason": "timeout",
      "error_message": "value",
      "content_url": "value",
      "content_mime_type": "value",
      "type": "image"
    }
  }
}`)
	}

	req := receiveWireFlowsTemplatesRuns(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/flows/templates/value/runs" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/flows/templates/value/runs")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsFlowsTemplatesRuns(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireFlowsTemplatesRuns(t, 422, `{
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

	var request templates.TemplateRunCreateRequest
	if err := json.Unmarshal([]byte(`{
  "template_id": "value",
  "inputs": {
    "key": "value"
  },
  "version_id": "value",
  "webhook": {
    "type": "all"
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Create(context.Background(), string("value"), &request)
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

	server, recorded := serveWireFlowsTemplatesRuns(t, 200, `{
  "id": "value",
  "template_id": "value",
  "version_id": "value",
  "status": "pending",
  "outputs": {
    "key": {
      "id": "value",
      "status": "pending",
      "failure_reason": "timeout",
      "error_message": "value",
      "content_url": "value",
      "content_mime_type": "value",
      "type": "image"
    }
  }
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
	if !wireJSONBytesEqualFlowsTemplatesRuns(gotResponse, []byte(`{
  "id": "value",
  "template_id": "value",
  "version_id": "value",
  "status": "pending",
  "outputs": {
    "key": {
      "id": "value",
      "status": "pending",
      "failure_reason": "timeout",
      "error_message": "value",
      "content_url": "value",
      "content_mime_type": "value",
      "type": "image"
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "template_id": "value",
  "version_id": "value",
  "status": "pending",
  "outputs": {
    "key": {
      "id": "value",
      "status": "pending",
      "failure_reason": "timeout",
      "error_message": "value",
      "content_url": "value",
      "content_mime_type": "value",
      "type": "image"
    }
  }
}`)
	}

	req := receiveWireFlowsTemplatesRuns(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/flows/templates/value/runs/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/flows/templates/value/runs/value2")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireFlowsTemplatesRuns(t, 422, `{
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

type wireRequestFlowsTemplatesRuns struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireFlowsTemplatesRuns records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireFlowsTemplatesRuns(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestFlowsTemplatesRuns) {
	t.Helper()
	recorded := make(chan wireRequestFlowsTemplatesRuns, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestFlowsTemplatesRuns{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireFlowsTemplatesRuns(t *testing.T, recorded <-chan wireRequestFlowsTemplatesRuns) wireRequestFlowsTemplatesRuns {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestFlowsTemplatesRuns{}
	}
}

// wireJSONContainsFlowsTemplatesRuns reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsFlowsTemplatesRuns(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualFlowsTemplatesRuns(got, want)
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

// setWireValueFlowsTemplatesRuns sets a named string, int or float field, or the value a pointer field points to.
func setWireValueFlowsTemplatesRuns(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualFlowsTemplatesRuns(got, want []byte) bool {
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
