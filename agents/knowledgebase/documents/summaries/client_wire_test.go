package summaries

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

	documents "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase/documents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsSummaries(t, 200, `{
  "key": {
    "status": "success",
    "data": {
      "id": "value",
      "name": "value",
      "metadata": {
        "created_at_unix_secs": 123,
        "last_updated_at_unix_secs": 123,
        "size_bytes": 123
      },
      "supported_usages": [
        "prompt"
      ],
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "folder_parent_id": "value",
      "chunking_override": {
        "max_chunk_size": 123,
        "overlap_percentage": 123,
        "custom_separators": [
          "value"
        ]
      },
      "folder_path": [
        {
          "id": "value"
        }
      ],
      "dependent_agents": [
        {
          "referenced_resource_ids": [
            "value"
          ],
          "id": "value",
          "name": "value",
          "type": "available",
          "created_at_unix_secs": 123,
          "access_level": "admin"
        }
      ],
      "type": "url",
      "url": "value",
      "auto_sync_info": {
        "minimum_frequency_days": 123,
        "auto_remove": true,
        "consec_failures": 123,
        "next_refresh_by": 123
      }
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &documents.GetSummariesRequest{}
	result, err := client.Get(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsSummaries(gotResponse, []byte(`{
  "key": {
    "status": "success",
    "data": {
      "id": "value",
      "name": "value",
      "metadata": {
        "created_at_unix_secs": 123,
        "last_updated_at_unix_secs": 123,
        "size_bytes": 123
      },
      "supported_usages": [
        "prompt"
      ],
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "folder_parent_id": "value",
      "chunking_override": {
        "max_chunk_size": 123,
        "overlap_percentage": 123,
        "custom_separators": [
          "value"
        ]
      },
      "folder_path": [
        {
          "id": "value"
        }
      ],
      "dependent_agents": [
        {
          "referenced_resource_ids": [
            "value"
          ],
          "id": "value",
          "name": "value",
          "type": "available",
          "created_at_unix_secs": 123,
          "access_level": "admin"
        }
      ],
      "type": "url",
      "url": "value",
      "auto_sync_info": {
        "minimum_frequency_days": 123,
        "auto_remove": true,
        "consec_failures": 123,
        "next_refresh_by": 123
      }
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": {
    "status": "success",
    "data": {
      "id": "value",
      "name": "value",
      "metadata": {
        "created_at_unix_secs": 123,
        "last_updated_at_unix_secs": 123,
        "size_bytes": 123
      },
      "supported_usages": [
        "prompt"
      ],
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "folder_parent_id": "value",
      "chunking_override": {
        "max_chunk_size": 123,
        "overlap_percentage": 123,
        "custom_separators": [
          "value"
        ]
      },
      "folder_path": [
        {
          "id": "value"
        }
      ],
      "dependent_agents": [
        {
          "referenced_resource_ids": [
            "value"
          ],
          "id": "value",
          "name": "value",
          "type": "available",
          "created_at_unix_secs": 123,
          "access_level": "admin"
        }
      ],
      "type": "url",
      "url": "value",
      "auto_sync_info": {
        "minimum_frequency_days": 123,
        "auto_remove": true,
        "consec_failures": 123,
        "next_refresh_by": 123
      }
    }
  }
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsSummaries(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/knowledge-base/summaries" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/summaries")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsSummaries(t, 422, `{
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

	request := &documents.GetSummariesRequest{}
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

type wireRequestAgentsKnowledgebaseDocumentsSummaries struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsKnowledgebaseDocumentsSummaries records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsKnowledgebaseDocumentsSummaries(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsKnowledgebaseDocumentsSummaries) {
	t.Helper()
	recorded := make(chan wireRequestAgentsKnowledgebaseDocumentsSummaries, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsKnowledgebaseDocumentsSummaries{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsKnowledgebaseDocumentsSummaries(t *testing.T, recorded <-chan wireRequestAgentsKnowledgebaseDocumentsSummaries) wireRequestAgentsKnowledgebaseDocumentsSummaries {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsKnowledgebaseDocumentsSummaries{}
	}
}

// wireJSONContainsAgentsKnowledgebaseDocumentsSummaries reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsKnowledgebaseDocumentsSummaries(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsKnowledgebaseDocumentsSummaries(got, want)
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

// setWireValueAgentsKnowledgebaseDocumentsSummaries sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsKnowledgebaseDocumentsSummaries(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsKnowledgebaseDocumentsSummaries(got, want []byte) bool {
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
