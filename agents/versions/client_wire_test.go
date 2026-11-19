package versions

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

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsVersions(t, 200, `{
  "id": "value",
  "agent_id": "value",
  "branch_id": "value",
  "version_description": "value",
  "seq_no_in_branch": 123,
  "time_committed_secs": 123,
  "parents": {
    "in_branch_parent_id": "value",
    "out_of_branch_parent_id": "value",
    "merged_into_branch_id": "value",
    "merged_from_branch_id": "value",
    "merged_from_version_id": "value",
    "rebased_from_version_id": "value"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "merged_authors": [
    {
      "is_creator": true,
      "creator_name": "value",
      "creator_email": "value",
      "role": "admin",
      "anonymous_access_level_override": "admin",
      "access_source": "creator"
    }
  ]
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
	if !wireJSONBytesEqualAgentsVersions(gotResponse, []byte(`{
  "id": "value",
  "agent_id": "value",
  "branch_id": "value",
  "version_description": "value",
  "seq_no_in_branch": 123,
  "time_committed_secs": 123,
  "parents": {
    "in_branch_parent_id": "value",
    "out_of_branch_parent_id": "value",
    "merged_into_branch_id": "value",
    "merged_from_branch_id": "value",
    "merged_from_version_id": "value",
    "rebased_from_version_id": "value"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "merged_authors": [
    {
      "is_creator": true,
      "creator_name": "value",
      "creator_email": "value",
      "role": "admin",
      "anonymous_access_level_override": "admin",
      "access_source": "creator"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "agent_id": "value",
  "branch_id": "value",
  "version_description": "value",
  "seq_no_in_branch": 123,
  "time_committed_secs": 123,
  "parents": {
    "in_branch_parent_id": "value",
    "out_of_branch_parent_id": "value",
    "merged_into_branch_id": "value",
    "merged_from_branch_id": "value",
    "merged_from_version_id": "value",
    "rebased_from_version_id": "value"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "merged_authors": [
    {
      "is_creator": true,
      "creator_name": "value",
      "creator_email": "value",
      "role": "admin",
      "anonymous_access_level_override": "admin",
      "access_source": "creator"
    }
  ]
}`)
	}

	req := receiveWireAgentsVersions(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agents/value/versions/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agents/value/versions/value2")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsVersions(t, 422, `{
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

type wireRequestAgentsVersions struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsVersions records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsVersions(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsVersions) {
	t.Helper()
	recorded := make(chan wireRequestAgentsVersions, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsVersions{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsVersions(t *testing.T, recorded <-chan wireRequestAgentsVersions) wireRequestAgentsVersions {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsVersions{}
	}
}

// wireJSONContainsAgentsVersions reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsVersions(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsVersions(got, want)
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

// setWireValueAgentsVersions sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsVersions(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsVersions(got, want []byte) bool {
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
