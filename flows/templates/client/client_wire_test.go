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
	flows "github.com/elevenlabs/elevenlabs-go/flows"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireFlowsTemplatesClient(t, 200, `{
  "templates": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "versions": [
        {
          "version_id": "value",
          "published_at_unix": 123,
          "is_latest": true,
          "inputs": [
            {
              "id": "value",
              "content_schema": {
                "title": "value",
                "description": "value",
                "type": "string",
                "enum": [
                  "value"
                ]
              }
            }
          ],
          "outputs": [
            {
              "id": "value",
              "content_schema": {
                "title": "value",
                "description": "value",
                "type": "string",
                "enum": [
                  "value"
                ]
              }
            }
          ]
        }
      ],
      "has_more_versions": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(25)
	versionsPerTemplateValue := int(10)
	searchValue := string("value")
	request := &flows.ListTemplatesRequest{
		Cursor:              &cursorValue,
		PageSize:            &pageSizeValue,
		VersionsPerTemplate: &versionsPerTemplateValue,
		Search:              &searchValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualFlowsTemplatesClient(gotResponse, []byte(`{
  "templates": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "versions": [
        {
          "version_id": "value",
          "published_at_unix": 123,
          "is_latest": true,
          "inputs": [
            {
              "id": "value",
              "content_schema": {
                "title": "value",
                "description": "value",
                "type": "string",
                "enum": [
                  "value"
                ]
              }
            }
          ],
          "outputs": [
            {
              "id": "value",
              "content_schema": {
                "title": "value",
                "description": "value",
                "type": "string",
                "enum": [
                  "value"
                ]
              }
            }
          ]
        }
      ],
      "has_more_versions": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "templates": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "versions": [
        {
          "version_id": "value",
          "published_at_unix": 123,
          "is_latest": true,
          "inputs": [
            {
              "id": "value",
              "content_schema": {
                "title": "value",
                "description": "value",
                "type": "string",
                "enum": [
                  "value"
                ]
              }
            }
          ],
          "outputs": [
            {
              "id": "value",
              "content_schema": {
                "title": "value",
                "description": "value",
                "type": "string",
                "enum": [
                  "value"
                ]
              }
            }
          ]
        }
      ],
      "has_more_versions": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireFlowsTemplatesClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/flows/templates" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/flows/templates")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "25" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "25")
	}
	if got := req.Query.Get("versions_per_template"); got != "10" {
		t.Fatalf("query %s: got %q want %q", "versions_per_template", got, "10")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireFlowsTemplatesClient(t, 422, `{
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
	pageSizeValue := int(25)
	versionsPerTemplateValue := int(10)
	searchValue := string("value")
	request := &flows.ListTemplatesRequest{
		Cursor:              &cursorValue,
		PageSize:            &pageSizeValue,
		VersionsPerTemplate: &versionsPerTemplateValue,
		Search:              &searchValue,
	}
	_, err := client.List(context.Background(), request)
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

	server, recorded := serveWireFlowsTemplatesClient(t, 200, `{
  "id": "value",
  "name": "value",
  "description": "value",
  "versions": [
    {
      "version_id": "value",
      "published_at_unix": 123,
      "is_latest": true,
      "inputs": [
        {
          "id": "value",
          "content_schema": {
            "title": "value",
            "description": "value",
            "type": "string",
            "enum": [
              "value"
            ]
          }
        }
      ],
      "outputs": [
        {
          "id": "value",
          "content_schema": {
            "title": "value",
            "description": "value",
            "type": "string",
            "enum": [
              "value"
            ]
          }
        }
      ]
    }
  ],
  "has_more_versions": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	versionsPerTemplateValue := int(10)
	request := &flows.GetTemplatesRequest{
		VersionsPerTemplate: &versionsPerTemplateValue,
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualFlowsTemplatesClient(gotResponse, []byte(`{
  "id": "value",
  "name": "value",
  "description": "value",
  "versions": [
    {
      "version_id": "value",
      "published_at_unix": 123,
      "is_latest": true,
      "inputs": [
        {
          "id": "value",
          "content_schema": {
            "title": "value",
            "description": "value",
            "type": "string",
            "enum": [
              "value"
            ]
          }
        }
      ],
      "outputs": [
        {
          "id": "value",
          "content_schema": {
            "title": "value",
            "description": "value",
            "type": "string",
            "enum": [
              "value"
            ]
          }
        }
      ]
    }
  ],
  "has_more_versions": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "name": "value",
  "description": "value",
  "versions": [
    {
      "version_id": "value",
      "published_at_unix": 123,
      "is_latest": true,
      "inputs": [
        {
          "id": "value",
          "content_schema": {
            "title": "value",
            "description": "value",
            "type": "string",
            "enum": [
              "value"
            ]
          }
        }
      ],
      "outputs": [
        {
          "id": "value",
          "content_schema": {
            "title": "value",
            "description": "value",
            "type": "string",
            "enum": [
              "value"
            ]
          }
        }
      ]
    }
  ],
  "has_more_versions": true
}`)
	}

	req := receiveWireFlowsTemplatesClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/flows/templates/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/flows/templates/value")
	}
	if got := req.Query.Get("versions_per_template"); got != "10" {
		t.Fatalf("query %s: got %q want %q", "versions_per_template", got, "10")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireFlowsTemplatesClient(t, 422, `{
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

	versionsPerTemplateValue := int(10)
	request := &flows.GetTemplatesRequest{
		VersionsPerTemplate: &versionsPerTemplateValue,
	}
	_, err := client.Get(context.Background(), string("value"), request)
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

type wireRequestFlowsTemplatesClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireFlowsTemplatesClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireFlowsTemplatesClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestFlowsTemplatesClient) {
	t.Helper()
	recorded := make(chan wireRequestFlowsTemplatesClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestFlowsTemplatesClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireFlowsTemplatesClient(t *testing.T, recorded <-chan wireRequestFlowsTemplatesClient) wireRequestFlowsTemplatesClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestFlowsTemplatesClient{}
	}
}

// wireJSONContainsFlowsTemplatesClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsFlowsTemplatesClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualFlowsTemplatesClient(got, want)
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

// setWireValueFlowsTemplatesClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueFlowsTemplatesClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualFlowsTemplatesClient(got, want []byte) bool {
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
