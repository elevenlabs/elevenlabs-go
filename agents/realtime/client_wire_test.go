package realtime

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

func TestListConversationGroupsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsRealtime(t, 200, `{
  "conversation_groups": [
    {
      "id": "value",
      "title": "value",
      "description": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "sharing": {
        "people_count": 123,
        "shared_with_everyone": true,
        "everyone_role": "admin"
      },
      "created_at_unix_secs": 123
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
	cursorValue := string("value")
	request := &agents.ListConversationGroupsRealtimeRequest{
		PageSize: &pageSizeValue,
		Search:   &searchValue,
		Cursor:   &cursorValue,
	}
	result, err := client.ListConversationGroups(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsRealtime(gotResponse, []byte(`{
  "conversation_groups": [
    {
      "id": "value",
      "title": "value",
      "description": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "sharing": {
        "people_count": 123,
        "shared_with_everyone": true,
        "everyone_role": "admin"
      },
      "created_at_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "conversation_groups": [
    {
      "id": "value",
      "title": "value",
      "description": "value",
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "sharing": {
        "people_count": 123,
        "shared_with_everyone": true,
        "everyone_role": "admin"
      },
      "created_at_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsRealtime(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/realtime/conversation-groups" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/realtime/conversation-groups")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListConversationGroupsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsRealtime(t, 422, `{
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
	cursorValue := string("value")
	request := &agents.ListConversationGroupsRealtimeRequest{
		PageSize: &pageSizeValue,
		Search:   &searchValue,
		Cursor:   &cursorValue,
	}
	_, err := client.ListConversationGroups(context.Background(), request)
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

func TestCreateConversationGroupWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsRealtime(t, 200, `{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.CreateRealtimeConversationGroupRequestModel
	if err := json.Unmarshal([]byte(`{
  "title": "value",
  "description": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CreateConversationGroup(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsRealtime(gotResponse, []byte(`{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`)
	}

	req := receiveWireAgentsRealtime(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/realtime/conversation-groups" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/realtime/conversation-groups")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsRealtime(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateConversationGroupWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsRealtime(t, 422, `{
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

	var request agents.CreateRealtimeConversationGroupRequestModel
	if err := json.Unmarshal([]byte(`{
  "title": "value",
  "description": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CreateConversationGroup(context.Background(), &request)
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

func TestGetConversationGroupWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsRealtime(t, 200, `{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.GetConversationGroup(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsRealtime(gotResponse, []byte(`{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`)
	}

	req := receiveWireAgentsRealtime(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/realtime/conversation-groups/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/realtime/conversation-groups/value")
	}
}

func TestGetConversationGroupWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsRealtime(t, 422, `{
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

	_, err := client.GetConversationGroup(context.Background(), string("value"))
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

func TestUpdateConversationGroupWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsRealtime(t, 200, `{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.UpdateRealtimeConversationGroupRequestModel
	if err := json.Unmarshal([]byte(`{
  "conversation_group_id": "value",
  "title": "value",
  "description": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.UpdateConversationGroup(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsRealtime(gotResponse, []byte(`{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "title": "value",
  "description": "value",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "sharing": {
    "people_count": 123,
    "shared_with_everyone": true,
    "everyone_role": "admin"
  },
  "created_at_unix_secs": 123
}`)
	}

	req := receiveWireAgentsRealtime(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/realtime/conversation-groups/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/realtime/conversation-groups/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsRealtime(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateConversationGroupWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsRealtime(t, 422, `{
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

	var request agents.UpdateRealtimeConversationGroupRequestModel
	if err := json.Unmarshal([]byte(`{
  "conversation_group_id": "value",
  "title": "value",
  "description": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.UpdateConversationGroup(context.Background(), string("value"), &request)
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

func TestDeleteConversationGroupWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsRealtime(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.DeleteConversationGroup(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsRealtime(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsRealtime(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/realtime/conversation-groups/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/realtime/conversation-groups/value")
	}
}

func TestDeleteConversationGroupWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsRealtime(t, 422, `{
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

	_, err := client.DeleteConversationGroup(context.Background(), string("value"))
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

func TestSetConversationGroupWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsRealtime(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.SetConversationGroupRequestModel
	if err := json.Unmarshal([]byte(`{
  "conversation_id": "value",
  "conversation_group_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.SetConversationGroup(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsRealtime(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsRealtime(t, recorded)
	if req.Method != "PUT" {
		t.Fatalf("method: got %q want %q", req.Method, "PUT")
	}
	if req.Path != "/v1/convai/realtime/conversations/value/conversation-group" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/realtime/conversations/value/conversation-group")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsRealtime(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestSetConversationGroupWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsRealtime(t, 422, `{
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

	var request agents.SetConversationGroupRequestModel
	if err := json.Unmarshal([]byte(`{
  "conversation_id": "value",
  "conversation_group_id": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.SetConversationGroup(context.Background(), string("value"), &request)
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

func TestClearConversationGroupWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsRealtime(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.ClearConversationGroup(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsRealtime(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsRealtime(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/realtime/conversations/value/conversation-group" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/realtime/conversations/value/conversation-group")
	}
}

func TestClearConversationGroupWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsRealtime(t, 422, `{
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

	_, err := client.ClearConversationGroup(context.Background(), string("value"))
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

func TestCreateClientSecretWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsRealtime(t, 200, `{
  "value": "value",
  "expires_at": 123,
  "session": {
    "type": "realtime",
    "model": "eleven_realtime_v1_mini",
    "instructions": "value",
    "output_modalities": [
      "value"
    ],
    "audio": {
      "input": {
        "format": {
          "type": "audio/pcm",
          "rate": 123
        },
        "transcription": {
          "language": "value"
        },
        "turn_detection": {
          "type": "semantic_vad",
          "eagerness": "low",
          "create_response": true,
          "interrupt_response": true,
          "threshold": 9.99,
          "prefix_padding_ms": 123,
          "silence_duration_ms": 123,
          "turn_timeout": 9.99
        },
        "pacing": "realtime"
      },
      "output": {
        "format": {
          "type": "audio/pcm"
        },
        "voice": "value"
      }
    },
    "tools": [
      {
        "type": "function",
        "name": "value",
        "description": "value",
        "parameters": {
          "key": {}
        }
      }
    ],
    "tool_choice": "value",
    "temperature": 9.99,
    "max_response_output_tokens": 123,
    "max_duration_seconds": 123,
    "tracing": "value"
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.CreateClientSecretRequest
	if err := json.Unmarshal([]byte(`{
  "session": {
    "type": "realtime",
    "model": "eleven_realtime_v1_mini",
    "instructions": "value",
    "output_modalities": [
      "value"
    ],
    "audio": {
      "input": {
        "format": {
          "type": "audio/pcm",
          "rate": 123
        },
        "transcription": {
          "language": "value"
        },
        "turn_detection": {
          "type": "semantic_vad",
          "eagerness": "low",
          "create_response": true,
          "interrupt_response": true,
          "threshold": 9.99,
          "prefix_padding_ms": 123,
          "silence_duration_ms": 123,
          "turn_timeout": 9.99
        },
        "pacing": "realtime"
      },
      "output": {
        "format": {
          "type": "audio/pcm"
        },
        "voice": "value"
      }
    },
    "tools": [
      {
        "type": "function",
        "name": "value",
        "description": "value",
        "parameters": {
          "key": {}
        }
      }
    ],
    "tool_choice": "value",
    "temperature": 9.99,
    "max_response_output_tokens": 123,
    "max_duration_seconds": 123,
    "turn_detection": {
      "type": "semantic_vad",
      "eagerness": "low",
      "create_response": true,
      "interrupt_response": true,
      "threshold": 9.99,
      "prefix_padding_ms": 123,
      "silence_duration_ms": 123,
      "turn_timeout": 9.99
    },
    "tracing": "value",
    "include": [
      "value"
    ],
    "parallel_tool_calls": true,
    "reasoning": {
      "key": {}
    },
    "truncation": "value",
    "prompt": {
      "key": {}
    }
  },
  "expires_after": {
    "anchor": "created_at",
    "seconds": 123
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.CreateClientSecret(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsRealtime(gotResponse, []byte(`{
  "value": "value",
  "expires_at": 123,
  "session": {
    "type": "realtime",
    "model": "eleven_realtime_v1_mini",
    "instructions": "value",
    "output_modalities": [
      "value"
    ],
    "audio": {
      "input": {
        "format": {
          "type": "audio/pcm",
          "rate": 123
        },
        "transcription": {
          "language": "value"
        },
        "turn_detection": {
          "type": "semantic_vad",
          "eagerness": "low",
          "create_response": true,
          "interrupt_response": true,
          "threshold": 9.99,
          "prefix_padding_ms": 123,
          "silence_duration_ms": 123,
          "turn_timeout": 9.99
        },
        "pacing": "realtime"
      },
      "output": {
        "format": {
          "type": "audio/pcm"
        },
        "voice": "value"
      }
    },
    "tools": [
      {
        "type": "function",
        "name": "value",
        "description": "value",
        "parameters": {
          "key": {}
        }
      }
    ],
    "tool_choice": "value",
    "temperature": 9.99,
    "max_response_output_tokens": 123,
    "max_duration_seconds": 123,
    "tracing": "value"
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "value": "value",
  "expires_at": 123,
  "session": {
    "type": "realtime",
    "model": "eleven_realtime_v1_mini",
    "instructions": "value",
    "output_modalities": [
      "value"
    ],
    "audio": {
      "input": {
        "format": {
          "type": "audio/pcm",
          "rate": 123
        },
        "transcription": {
          "language": "value"
        },
        "turn_detection": {
          "type": "semantic_vad",
          "eagerness": "low",
          "create_response": true,
          "interrupt_response": true,
          "threshold": 9.99,
          "prefix_padding_ms": 123,
          "silence_duration_ms": 123,
          "turn_timeout": 9.99
        },
        "pacing": "realtime"
      },
      "output": {
        "format": {
          "type": "audio/pcm"
        },
        "voice": "value"
      }
    },
    "tools": [
      {
        "type": "function",
        "name": "value",
        "description": "value",
        "parameters": {
          "key": {}
        }
      }
    ],
    "tool_choice": "value",
    "temperature": 9.99,
    "max_response_output_tokens": 123,
    "max_duration_seconds": 123,
    "tracing": "value"
  }
}`)
	}

	req := receiveWireAgentsRealtime(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/realtime/client_secrets" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/realtime/client_secrets")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsRealtime(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateClientSecretWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsRealtime(t, 422, `{
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

	var request agents.CreateClientSecretRequest
	if err := json.Unmarshal([]byte(`{
  "session": {
    "type": "realtime",
    "model": "eleven_realtime_v1_mini",
    "instructions": "value",
    "output_modalities": [
      "value"
    ],
    "audio": {
      "input": {
        "format": {
          "type": "audio/pcm",
          "rate": 123
        },
        "transcription": {
          "language": "value"
        },
        "turn_detection": {
          "type": "semantic_vad",
          "eagerness": "low",
          "create_response": true,
          "interrupt_response": true,
          "threshold": 9.99,
          "prefix_padding_ms": 123,
          "silence_duration_ms": 123,
          "turn_timeout": 9.99
        },
        "pacing": "realtime"
      },
      "output": {
        "format": {
          "type": "audio/pcm"
        },
        "voice": "value"
      }
    },
    "tools": [
      {
        "type": "function",
        "name": "value",
        "description": "value",
        "parameters": {
          "key": {}
        }
      }
    ],
    "tool_choice": "value",
    "temperature": 9.99,
    "max_response_output_tokens": 123,
    "max_duration_seconds": 123,
    "turn_detection": {
      "type": "semantic_vad",
      "eagerness": "low",
      "create_response": true,
      "interrupt_response": true,
      "threshold": 9.99,
      "prefix_padding_ms": 123,
      "silence_duration_ms": 123,
      "turn_timeout": 9.99
    },
    "tracing": "value",
    "include": [
      "value"
    ],
    "parallel_tool_calls": true,
    "reasoning": {
      "key": {}
    },
    "truncation": "value",
    "prompt": {
      "key": {}
    }
  },
  "expires_after": {
    "anchor": "created_at",
    "seconds": 123
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.CreateClientSecret(context.Background(), &request)
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

type wireRequestAgentsRealtime struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsRealtime records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsRealtime(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsRealtime) {
	t.Helper()
	recorded := make(chan wireRequestAgentsRealtime, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsRealtime{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsRealtime(t *testing.T, recorded <-chan wireRequestAgentsRealtime) wireRequestAgentsRealtime {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsRealtime{}
	}
}

// wireJSONContainsAgentsRealtime reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsRealtime(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsRealtime(got, want)
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

// setWireValueAgentsRealtime sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsRealtime(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsRealtime(got, want []byte) bool {
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
