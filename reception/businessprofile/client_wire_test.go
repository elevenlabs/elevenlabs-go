package businessprofile

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
	reception "github.com/elevenlabs/elevenlabs-go/reception"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionBusinessprofile(t, 200, `{
  "business_name": "value",
  "business_country": "value",
  "business_description": "value",
  "timezone": "value",
  "currency": "USD",
  "operating_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "tool_call_sound": "typing",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "workspace_now": {
    "today": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "hour": 123,
    "minute": 123,
    "fetched_at_ms": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.GetBusinessProfileRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.Get(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionBusinessprofile(gotResponse, []byte(`{
  "business_name": "value",
  "business_country": "value",
  "business_description": "value",
  "timezone": "value",
  "currency": "USD",
  "operating_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "tool_call_sound": "typing",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "workspace_now": {
    "today": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "hour": 123,
    "minute": 123,
    "fetched_at_ms": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "business_name": "value",
  "business_country": "value",
  "business_description": "value",
  "timezone": "value",
  "currency": "USD",
  "operating_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "tool_call_sound": "typing",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "workspace_now": {
    "today": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "hour": 123,
    "minute": 123,
    "fetched_at_ms": 123
  }
}`)
	}

	req := receiveWireReceptionBusinessprofile(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/agent-info" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/agent-info")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionBusinessprofile(t, 422, `{
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

	request := &reception.GetBusinessProfileRequest{
		XiWorkspaceId: "value",
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionBusinessprofile(t, 200, `{
  "business_name": "value",
  "business_country": "value",
  "business_description": "value",
  "timezone": "value",
  "currency": "USD",
  "operating_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "tool_call_sound": "typing",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "workspace_now": {
    "today": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "hour": 123,
    "minute": 123,
    "fetched_at_ms": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReceptionUpdateAgentInfoRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "business_name": "value",
  "business_country": "value",
  "business_description": "value",
  "timezone": "value",
  "currency": "USD",
  "operating_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "tool_call_sound": "typing"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Update(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionBusinessprofile(gotResponse, []byte(`{
  "business_name": "value",
  "business_country": "value",
  "business_description": "value",
  "timezone": "value",
  "currency": "USD",
  "operating_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "tool_call_sound": "typing",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "workspace_now": {
    "today": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "hour": 123,
    "minute": 123,
    "fetched_at_ms": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "business_name": "value",
  "business_country": "value",
  "business_description": "value",
  "timezone": "value",
  "currency": "USD",
  "operating_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "tool_call_sound": "typing",
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "workspace_now": {
    "today": {
      "year": 2026,
      "month": 12,
      "day": 31
    },
    "hour": 123,
    "minute": 123,
    "fetched_at_ms": 123
  }
}`)
	}

	req := receiveWireReceptionBusinessprofile(t, recorded)
	if req.Method != "PUT" {
		t.Fatalf("method: got %q want %q", req.Method, "PUT")
	}
	if req.Path != "/v1/reception/agent-info" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/agent-info")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionBusinessprofile(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionBusinessprofile(t, 422, `{
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

	var request reception.ReceptionUpdateAgentInfoRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "business_name": "value",
  "business_country": "value",
  "business_description": "value",
  "timezone": "value",
  "currency": "USD",
  "operating_hours": {
    "hours": {
      "key": [
        {
          "start_hour": 123,
          "start_minute": 123,
          "end_hour": 24,
          "end_minute": 123
        }
      ]
    },
    "timezone": "value"
  },
  "tool_call_sound": "typing"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	_, err := client.Update(context.Background(), &request)
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

type wireRequestReceptionBusinessprofile struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionBusinessprofile records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionBusinessprofile(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionBusinessprofile) {
	t.Helper()
	recorded := make(chan wireRequestReceptionBusinessprofile, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionBusinessprofile{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionBusinessprofile(t *testing.T, recorded <-chan wireRequestReceptionBusinessprofile) wireRequestReceptionBusinessprofile {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionBusinessprofile{}
	}
}

// wireJSONContainsReceptionBusinessprofile reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionBusinessprofile(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionBusinessprofile(got, want)
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

// setWireValueReceptionBusinessprofile sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionBusinessprofile(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionBusinessprofile(got, want []byte) bool {
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
