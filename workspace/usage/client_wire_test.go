package usage

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
	workspace "github.com/elevenlabs/elevenlabs-go/workspace"
)

func TestGetUsageByProductOverTimeWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceUsage(t, 200, `{
  "columns": [
    "value"
  ],
  "column_types": [
    "String"
  ],
  "rows": [
    [
      "value"
    ]
  ],
  "column_units": [
    "ms"
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request workspace.BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost
	if err := json.Unmarshal([]byte(`{
  "start_time": 1577836800000,
  "end_time": 1577836800000,
  "interval_seconds": 123,
  "group_by": [
    "product_type"
  ],
  "filters": [
    {
      "column": "value",
      "operation": "in",
      "values": [
        "value"
      ]
    }
  ],
  "time_zone": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.GetUsageByProductOverTime(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceUsage(gotResponse, []byte(`{
  "columns": [
    "value"
  ],
  "column_types": [
    "String"
  ],
  "rows": [
    [
      "value"
    ]
  ],
  "column_units": [
    "ms"
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "columns": [
    "value"
  ],
  "column_types": [
    "String"
  ],
  "rows": [
    [
      "value"
    ]
  ],
  "column_units": [
    "ms"
  ]
}`)
	}

	req := receiveWireWorkspaceUsage(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/analytics/query/usage-by-product-over-time" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/analytics/query/usage-by-product-over-time")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWorkspaceUsage(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestGetUsageByProductOverTimeWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceUsage(t, 422, `{
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

	var request workspace.BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost
	if err := json.Unmarshal([]byte(`{
  "start_time": 1577836800000,
  "end_time": 1577836800000,
  "interval_seconds": 123,
  "group_by": [
    "product_type"
  ],
  "filters": [
    {
      "column": "value",
      "operation": "in",
      "values": [
        "value"
      ]
    }
  ],
  "time_zone": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.GetUsageByProductOverTime(context.Background(), &request)
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

func TestGetUsageWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceUsage(t, 200, `{
  "columns": [
    "value"
  ],
  "column_types": [
    "String"
  ],
  "rows": [
    [
      "value"
    ]
  ],
  "column_units": [
    "ms"
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request workspace.BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost
	if err := json.Unmarshal([]byte(`{
  "start_time": 1577836800000,
  "end_time": 1577836800000,
  "interval_seconds": 86400,
  "group_by": [
    "model"
  ],
  "filters": [
    {
      "column": "model",
      "operation": "eq",
      "values": [
        "value"
      ]
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.GetUsage(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceUsage(gotResponse, []byte(`{
  "columns": [
    "value"
  ],
  "column_types": [
    "String"
  ],
  "rows": [
    [
      "value"
    ]
  ],
  "column_units": [
    "ms"
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "columns": [
    "value"
  ],
  "column_types": [
    "String"
  ],
  "rows": [
    [
      "value"
    ]
  ],
  "column_units": [
    "ms"
  ]
}`)
	}

	req := receiveWireWorkspaceUsage(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/workspace/analytics/usage" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/analytics/usage")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsWorkspaceUsage(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestGetUsageWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceUsage(t, 422, `{
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

	var request workspace.BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost
	if err := json.Unmarshal([]byte(`{
  "start_time": 1577836800000,
  "end_time": 1577836800000,
  "interval_seconds": 86400,
  "group_by": [
    "model"
  ],
  "filters": [
    {
      "column": "model",
      "operation": "eq",
      "values": [
        "value"
      ]
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.GetUsage(context.Background(), &request)
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

type wireRequestWorkspaceUsage struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireWorkspaceUsage records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireWorkspaceUsage(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestWorkspaceUsage) {
	t.Helper()
	recorded := make(chan wireRequestWorkspaceUsage, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestWorkspaceUsage{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireWorkspaceUsage(t *testing.T, recorded <-chan wireRequestWorkspaceUsage) wireRequestWorkspaceUsage {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestWorkspaceUsage{}
	}
}

// wireJSONContainsWorkspaceUsage reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsWorkspaceUsage(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualWorkspaceUsage(got, want)
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

// setWireValueWorkspaceUsage sets a named string, int or float field, or the value a pointer field points to.
func setWireValueWorkspaceUsage(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualWorkspaceUsage(got, want []byte) bool {
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
