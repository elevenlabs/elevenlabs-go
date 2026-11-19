package auditlogs

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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireWorkspaceAuditlogs(t, 200, `{
  "entries": [
    {
      "metadata": {
        "key": {}
      },
      "time": 123,
      "activity_id": 123,
      "activity_name": "value",
      "category_name": "value",
      "category_uid": 123,
      "class_name": "value",
      "class_uid": 123,
      "severity_id": 123,
      "status_id": 123,
      "actor": {
        "user": {
          "name": "value",
          "uid": "value",
          "type_id": 123,
          "type": "value",
          "email_addr": "value",
          "full_name": "value",
          "domain": "value"
        },
        "app_name": "value",
        "app_uid": "value",
        "session": {
          "key": {}
        }
      },
      "device": {
        "ip": "value",
        "hostname": "value",
        "type_id": 123
      },
      "http_request": {
        "http_method": "value",
        "url": {
          "url_string": "value",
          "scheme": "value",
          "hostname": "value",
          "port": 123,
          "path": "value",
          "query_string": "value"
        },
        "user_agent": "value",
        "x_forwarded_for": [
          "value"
        ]
      },
      "message": "value",
      "unmapped": {
        "key": {}
      },
      "id": "value",
      "time_dt": "value",
      "type_uid": 123,
      "type_name": "value"
    }
  ],
  "has_more": true,
  "next_cursor": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	limitValue := int(100)
	cursorValue := string("value")
	timeFromUnixMsValue := int(123)
	timeToUnixMsValue := int(123)
	actorUidValue := string("value")
	request := &workspace.ListAuditLogsRequest{
		Limit:          &limitValue,
		Cursor:         &cursorValue,
		TimeFromUnixMs: &timeFromUnixMsValue,
		TimeToUnixMs:   &timeToUnixMsValue,
		ActorUid:       &actorUidValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualWorkspaceAuditlogs(gotResponse, []byte(`{
  "entries": [
    {
      "metadata": {
        "key": {}
      },
      "time": 123,
      "activity_id": 123,
      "activity_name": "value",
      "category_name": "value",
      "category_uid": 123,
      "class_name": "value",
      "class_uid": 123,
      "severity_id": 123,
      "status_id": 123,
      "actor": {
        "user": {
          "name": "value",
          "uid": "value",
          "type_id": 123,
          "type": "value",
          "email_addr": "value",
          "full_name": "value",
          "domain": "value"
        },
        "app_name": "value",
        "app_uid": "value",
        "session": {
          "key": {}
        }
      },
      "device": {
        "ip": "value",
        "hostname": "value",
        "type_id": 123
      },
      "http_request": {
        "http_method": "value",
        "url": {
          "url_string": "value",
          "scheme": "value",
          "hostname": "value",
          "port": 123,
          "path": "value",
          "query_string": "value"
        },
        "user_agent": "value",
        "x_forwarded_for": [
          "value"
        ]
      },
      "message": "value",
      "unmapped": {
        "key": {}
      },
      "id": "value",
      "time_dt": "value",
      "type_uid": 123,
      "type_name": "value"
    }
  ],
  "has_more": true,
  "next_cursor": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "entries": [
    {
      "metadata": {
        "key": {}
      },
      "time": 123,
      "activity_id": 123,
      "activity_name": "value",
      "category_name": "value",
      "category_uid": 123,
      "class_name": "value",
      "class_uid": 123,
      "severity_id": 123,
      "status_id": 123,
      "actor": {
        "user": {
          "name": "value",
          "uid": "value",
          "type_id": 123,
          "type": "value",
          "email_addr": "value",
          "full_name": "value",
          "domain": "value"
        },
        "app_name": "value",
        "app_uid": "value",
        "session": {
          "key": {}
        }
      },
      "device": {
        "ip": "value",
        "hostname": "value",
        "type_id": 123
      },
      "http_request": {
        "http_method": "value",
        "url": {
          "url_string": "value",
          "scheme": "value",
          "hostname": "value",
          "port": 123,
          "path": "value",
          "query_string": "value"
        },
        "user_agent": "value",
        "x_forwarded_for": [
          "value"
        ]
      },
      "message": "value",
      "unmapped": {
        "key": {}
      },
      "id": "value",
      "time_dt": "value",
      "type_uid": 123,
      "type_name": "value"
    }
  ],
  "has_more": true,
  "next_cursor": "value"
}`)
	}

	req := receiveWireWorkspaceAuditlogs(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/workspace/audit-logs" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/workspace/audit-logs")
	}
	if got := req.Query.Get("limit"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "limit", got, "100")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("time_from_unix_ms"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "time_from_unix_ms", got, "123")
	}
	if got := req.Query.Get("time_to_unix_ms"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "time_to_unix_ms", got, "123")
	}
	if got := req.Query.Get("actor_uid"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "actor_uid", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireWorkspaceAuditlogs(t, 422, `{
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

	limitValue := int(100)
	cursorValue := string("value")
	timeFromUnixMsValue := int(123)
	timeToUnixMsValue := int(123)
	actorUidValue := string("value")
	request := &workspace.ListAuditLogsRequest{
		Limit:          &limitValue,
		Cursor:         &cursorValue,
		TimeFromUnixMs: &timeFromUnixMsValue,
		TimeToUnixMs:   &timeToUnixMsValue,
		ActorUid:       &actorUidValue,
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

type wireRequestWorkspaceAuditlogs struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireWorkspaceAuditlogs records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireWorkspaceAuditlogs(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestWorkspaceAuditlogs) {
	t.Helper()
	recorded := make(chan wireRequestWorkspaceAuditlogs, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestWorkspaceAuditlogs{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireWorkspaceAuditlogs(t *testing.T, recorded <-chan wireRequestWorkspaceAuditlogs) wireRequestWorkspaceAuditlogs {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestWorkspaceAuditlogs{}
	}
}

// wireJSONContainsWorkspaceAuditlogs reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsWorkspaceAuditlogs(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualWorkspaceAuditlogs(got, want)
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

// setWireValueWorkspaceAuditlogs sets a named string, int or float field, or the value a pointer field points to.
func setWireValueWorkspaceAuditlogs(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualWorkspaceAuditlogs(got, want []byte) bool {
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
