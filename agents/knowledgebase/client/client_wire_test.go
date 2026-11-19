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

	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestSizeWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseClient(t, 200, `{
  "number_of_pages": 9.99
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Size(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseClient(gotResponse, []byte(`{
  "number_of_pages": 9.99
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "number_of_pages": 9.99
}`)
	}

	req := receiveWireAgentsKnowledgebaseClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/agent/value/knowledge-base/size" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/agent/value/knowledge-base/size")
	}
}

func TestSizeWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseClient(t, 422, `{
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

	_, err := client.Size(context.Background(), string("value"))
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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseClient(t, 200, `{
  "documents": [
    {
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
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	searchValue := string("value")
	showOnlyOwnedDocumentsValue := bool(true)
	createdByUserIdValue := string("value")
	parentFolderIdValue := string("value")
	ancestorFolderIdValue := string("value")
	foldersFirstValue := bool(true)
	cursorValue := string("value")
	request := &agents.ListKnowledgeBaseRequest{
		PageSize:               &pageSizeValue,
		Search:                 &searchValue,
		ShowOnlyOwnedDocuments: &showOnlyOwnedDocumentsValue,
		CreatedByUserId:        &createdByUserIdValue,
		ParentFolderId:         &parentFolderIdValue,
		AncestorFolderId:       &ancestorFolderIdValue,
		FoldersFirst:           &foldersFirstValue,
		Cursor:                 &cursorValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseClient(gotResponse, []byte(`{
  "documents": [
    {
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
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "documents": [
    {
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
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsKnowledgebaseClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/knowledge-base" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("show_only_owned_documents"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "show_only_owned_documents", got, "true")
	}
	if got := req.Query.Get("created_by_user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "created_by_user_id", got, "value")
	}
	if got := req.Query.Get("parent_folder_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "parent_folder_id", got, "value")
	}
	if got := req.Query.Get("ancestor_folder_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "ancestor_folder_id", got, "value")
	}
	if got := req.Query.Get("folders_first"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "folders_first", got, "true")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseClient(t, 422, `{
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

	pageSizeValue := int(100)
	searchValue := string("value")
	showOnlyOwnedDocumentsValue := bool(true)
	createdByUserIdValue := string("value")
	parentFolderIdValue := string("value")
	ancestorFolderIdValue := string("value")
	foldersFirstValue := bool(true)
	cursorValue := string("value")
	request := &agents.ListKnowledgeBaseRequest{
		PageSize:               &pageSizeValue,
		Search:                 &searchValue,
		ShowOnlyOwnedDocuments: &showOnlyOwnedDocumentsValue,
		CreatedByUserId:        &createdByUserIdValue,
		ParentFolderId:         &parentFolderIdValue,
		AncestorFolderId:       &ancestorFolderIdValue,
		FoldersFirst:           &foldersFirstValue,
		Cursor:                 &cursorValue,
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

func TestGetOrCreateRagIndexesWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseClient(t, 200, `{
  "key": {
    "status": "success",
    "data": {
      "id": "value",
      "model": "e5_mistral_7b_instruct",
      "status": "new",
      "progress_percentage": 9.99,
      "document_model_index_usage": {
        "used_bytes": 123
      }
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost
	if err := json.Unmarshal([]byte(`{
  "items": [
    {
      "document_id": "value",
      "create_if_missing": true,
      "model": "e5_mistral_7b_instruct"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.GetOrCreateRagIndexes(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseClient(gotResponse, []byte(`{
  "key": {
    "status": "success",
    "data": {
      "id": "value",
      "model": "e5_mistral_7b_instruct",
      "status": "new",
      "progress_percentage": 9.99,
      "document_model_index_usage": {
        "used_bytes": 123
      }
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": {
    "status": "success",
    "data": {
      "id": "value",
      "model": "e5_mistral_7b_instruct",
      "status": "new",
      "progress_percentage": 9.99,
      "document_model_index_usage": {
        "used_bytes": 123
      }
    }
  }
}`)
	}

	req := receiveWireAgentsKnowledgebaseClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/rag-index" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/rag-index")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestGetOrCreateRagIndexesWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseClient(t, 422, `{
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

	var request agents.BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost
	if err := json.Unmarshal([]byte(`{
  "items": [
    {
      "document_id": "value",
      "create_if_missing": true,
      "model": "e5_mistral_7b_instruct"
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.GetOrCreateRagIndexes(context.Background(), &request)
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

func TestSearchWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseClient(t, 200, `{
  "results": [
    {
      "document": {
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
      },
      "search_snippet": [
        {
          "value": "value",
          "is_hit": true
        }
      ],
      "score": 9.99
    }
  ],
  "next_cursor": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &agents.SearchKnowledgeBaseRequest{
		Query:    "value",
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	result, err := client.Search(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseClient(gotResponse, []byte(`{
  "results": [
    {
      "document": {
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
      },
      "search_snippet": [
        {
          "value": "value",
          "is_hit": true
        }
      ],
      "score": 9.99
    }
  ],
  "next_cursor": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "results": [
    {
      "document": {
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
      },
      "search_snippet": [
        {
          "value": "value",
          "is_hit": true
        }
      ],
      "score": 9.99
    }
  ],
  "next_cursor": "value"
}`)
	}

	req := receiveWireAgentsKnowledgebaseClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/knowledge-base/search" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/search")
	}
	if got := req.Query.Get("query"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "query", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestSearchWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseClient(t, 422, `{
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

	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &agents.SearchKnowledgeBaseRequest{
		Query:    "value",
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	_, err := client.Search(context.Background(), request)
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

type wireRequestAgentsKnowledgebaseClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsKnowledgebaseClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsKnowledgebaseClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsKnowledgebaseClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsKnowledgebaseClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsKnowledgebaseClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsKnowledgebaseClient(t *testing.T, recorded <-chan wireRequestAgentsKnowledgebaseClient) wireRequestAgentsKnowledgebaseClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsKnowledgebaseClient{}
	}
}

// wireJSONContainsAgentsKnowledgebaseClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsKnowledgebaseClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsKnowledgebaseClient(got, want)
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

// setWireValueAgentsKnowledgebaseClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsKnowledgebaseClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsKnowledgebaseClient(got, want []byte) bool {
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
