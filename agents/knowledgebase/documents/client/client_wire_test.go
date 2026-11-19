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

	knowledgebase "github.com/elevenlabs/elevenlabs-go/agents/knowledgebase"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestCreateFromUrlWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost
	if err := json.Unmarshal([]byte(`{
  "url": "value",
  "name": "value",
  "parent_folder_id": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "minimum_frequency_days": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CreateFromUrl(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/url" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/url")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseDocumentsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateFromUrlWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	var request knowledgebase.BodyCreateUrlDocumentV1ConvaiKnowledgeBaseUrlPost
	if err := json.Unmarshal([]byte(`{
  "url": "value",
  "name": "value",
  "parent_folder_id": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "minimum_frequency_days": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CreateFromUrl(context.Background(), &request)
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

func TestCreateFromTextWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost
	if err := json.Unmarshal([]byte(`{
  "text": "value",
  "name": "value",
  "parent_folder_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CreateFromText(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/text" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/text")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseDocumentsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateFromTextWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	var request knowledgebase.BodyCreateTextDocumentV1ConvaiKnowledgeBaseTextPost
	if err := json.Unmarshal([]byte(`{
  "text": "value",
  "name": "value",
  "parent_folder_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CreateFromText(context.Background(), &request)
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

func TestCreateFolderWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "parent_folder_id": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "minimum_frequency_days": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CreateFolder(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/folder" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/folder")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseDocumentsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateFolderWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	var request knowledgebase.BodyCreateFolderV1ConvaiKnowledgeBaseFolderPost
	if err := json.Unmarshal([]byte(`{
  "name": "value",
  "parent_folder_id": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "minimum_frequency_days": 123
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CreateFolder(context.Background(), &request)
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

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
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
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	agentIdValue := string("value")
	request := &knowledgebase.GetDocumentsRequest{
		AgentId: &agentIdValue,
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
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
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/knowledge-base/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/value")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	agentIdValue := string("value")
	request := &knowledgebase.GetDocumentsRequest{
		AgentId: &agentIdValue,
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
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
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch
	if err := json.Unmarshal([]byte(`{
  "documentation_id": "value",
  "name": "value",
  "content": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "minimum_frequency_days": 123,
  "tighten_only": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
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
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
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
      "id": "value",
      "name": "value"
    }
  ],
  "type": "url",
  "url": "value",
  "extracted_inner_html": "value",
  "content_format": "html",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/knowledge-base/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseDocumentsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	var request knowledgebase.BodyUpdateDocumentV1ConvaiKnowledgeBaseDocumentationIdPatch
	if err := json.Unmarshal([]byte(`{
  "documentation_id": "value",
  "name": "value",
  "content": "value",
  "enable_auto_sync": true,
  "auto_remove": true,
  "minimum_frequency_days": 123,
  "tighten_only": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), &request)
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

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	forceValue := bool(true)
	request := &knowledgebase.DeleteDocumentsRequest{
		Force: &forceValue,
	}
	result, err := client.Delete(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/knowledge-base/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/value")
	}
	if got := req.Query.Get("force"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "force", got, "true")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	forceValue := bool(true)
	request := &knowledgebase.DeleteDocumentsRequest{
		Force: &forceValue,
	}
	_, err := client.Delete(context.Background(), string("value"), request)
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

func TestGetChunkingOverrideWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
  "id": "value",
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "effective_chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "inherited_chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetChunkingOverride(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
  "id": "value",
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "effective_chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "inherited_chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "effective_chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "inherited_chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  }
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/knowledge-base/value/chunking-override" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/value/chunking-override")
	}
}

func TestGetChunkingOverrideWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	_, err := client.GetChunkingOverride(context.Background(), string("value"))
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

func TestUpdateChunkingOverrideWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
  "id": "value",
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "reindex_queued": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.UpdateDocumentChunkingOverrideRequestModel
	if err := json.Unmarshal([]byte(`{
  "documentation_id": "value",
  "chunking_override": {
    "max_chunk_size": 200,
    "overlap_percentage": 25,
    "custom_separators": [
      "value"
    ]
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.UpdateChunkingOverride(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
  "id": "value",
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "reindex_queued": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "reindex_queued": true
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/knowledge-base/value/chunking-override" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/value/chunking-override")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseDocumentsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateChunkingOverrideWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	var request knowledgebase.UpdateDocumentChunkingOverrideRequestModel
	if err := json.Unmarshal([]byte(`{
  "documentation_id": "value",
  "chunking_override": {
    "max_chunk_size": 200,
    "overlap_percentage": 25,
    "custom_separators": [
      "value"
    ]
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.UpdateChunkingOverride(context.Background(), string("value"), &request)
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

func TestGetAgentsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
  "agents": [
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
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(100)
	cursorValue := string("value")
	request := &knowledgebase.GetAgentsDocumentsRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	result, err := client.GetAgents(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
  "agents": [
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
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agents": [
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
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/knowledge-base/value/dependent-agents" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/value/dependent-agents")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestGetAgentsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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
	request := &knowledgebase.GetAgentsDocumentsRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	_, err := client.GetAgents(context.Background(), string("value"), request)
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

func TestGetBulkAgentsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
  "agents": [
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
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost
	if err := json.Unmarshal([]byte(`{
  "cursor": "value",
  "page_size": 100,
  "dependent_type": "direct",
  "document_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	pageSizeValue := int(100)
	request.PageSize = &pageSizeValue
	cursorValue := string("value")
	request.Cursor = &cursorValue
	result, err := client.GetBulkAgents(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
  "agents": [
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
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agents": [
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
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/dependent-agents" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/dependent-agents")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseDocumentsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestGetBulkAgentsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	var request knowledgebase.BodyGetDependentAgentsForMultipleDocumentsV1ConvaiKnowledgeBaseDependentAgentsPost
	if err := json.Unmarshal([]byte(`{
  "cursor": "value",
  "page_size": 100,
  "dependent_type": "direct",
  "document_ids": [
    "value"
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	pageSizeValue := int(100)
	request.PageSize = &pageSizeValue
	cursorValue := string("value")
	request.Cursor = &cursorValue
	_, err := client.GetBulkAgents(context.Background(), &request)
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

func TestGetSourceFileUrlWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
  "signed_url": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetSourceFileUrl(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
  "signed_url": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "signed_url": "value"
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/knowledge-base/value/source-file-url" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/value/source-file-url")
	}
}

func TestGetSourceFileUrlWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	_, err := client.GetSourceFileUrl(context.Background(), string("value"))
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

func TestMoveWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost
	if err := json.Unmarshal([]byte(`{
  "document_id": "value",
  "move_to": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	err := client.Move(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/value/move" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/value/move")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseDocumentsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestMoveWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	var request knowledgebase.BodyMoveEntityToFolderV1ConvaiKnowledgeBaseDocumentIdMovePost
	if err := json.Unmarshal([]byte(`{
  "document_id": "value",
  "move_to": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	err := client.Move(context.Background(), string("value"), &request)
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

func TestBulkMoveWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, "")

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost
	if err := json.Unmarshal([]byte(`{
  "document_ids": [
    "value"
  ],
  "move_to": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	err := client.BulkMove(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/bulk-move" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/bulk-move")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseDocumentsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestBulkMoveWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	var request knowledgebase.BodyBulkMoveEntitiesToFolderV1ConvaiKnowledgeBaseBulkMovePost
	if err := json.Unmarshal([]byte(`{
  "document_ids": [
    "value"
  ],
  "move_to": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	err := client.BulkMove(context.Background(), &request)
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

func TestBulkDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsKnowledgebaseDocumentsClient(t, 200, `{
  "key": {
    "status": "success",
    "data": {
      "id": "value"
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request knowledgebase.BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost
	if err := json.Unmarshal([]byte(`{
  "document_ids": [
    "value"
  ],
  "force": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.BulkDelete(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(gotResponse, []byte(`{
  "key": {
    "status": "success",
    "data": {
      "id": "value"
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "key": {
    "status": "success",
    "data": {
      "id": "value"
    }
  }
}`)
	}

	req := receiveWireAgentsKnowledgebaseDocumentsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/knowledge-base/bulk-delete" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/knowledge-base/bulk-delete")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsKnowledgebaseDocumentsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestBulkDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsKnowledgebaseDocumentsClient(t, 422, `{
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

	var request knowledgebase.BodyBulkDeleteKnowledgeBaseDocumentsV1ConvaiKnowledgeBaseBulkDeletePost
	if err := json.Unmarshal([]byte(`{
  "document_ids": [
    "value"
  ],
  "force": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.BulkDelete(context.Background(), &request)
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

type wireRequestAgentsKnowledgebaseDocumentsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsKnowledgebaseDocumentsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsKnowledgebaseDocumentsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsKnowledgebaseDocumentsClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsKnowledgebaseDocumentsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsKnowledgebaseDocumentsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsKnowledgebaseDocumentsClient(t *testing.T, recorded <-chan wireRequestAgentsKnowledgebaseDocumentsClient) wireRequestAgentsKnowledgebaseDocumentsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsKnowledgebaseDocumentsClient{}
	}
}

// wireJSONContainsAgentsKnowledgebaseDocumentsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsKnowledgebaseDocumentsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(got, want)
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

// setWireValueAgentsKnowledgebaseDocumentsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsKnowledgebaseDocumentsClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsKnowledgebaseDocumentsClient(got, want []byte) bool {
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
