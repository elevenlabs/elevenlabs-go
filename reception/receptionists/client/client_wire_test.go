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
	option "github.com/elevenlabs/elevenlabs-go/option"
	reception "github.com/elevenlabs/elevenlabs-go/reception"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionReceptionistsClient(t, 200, `[
  {
    "agent_id": "value",
    "config": {
      "additional_instructions": "value",
      "tone_traits": [
        "value"
      ],
      "details_to_collect": [
        "value"
      ],
      "llm": "gpt-4o-mini",
      "reasoning_effort": "none",
      "disable_rag": true,
      "background_voice_detection": true,
      "turn_eagerness": "patient",
      "voice_id": "value",
      "language": "value",
      "additional_languages": [
        "value"
      ],
      "first_message": "value",
      "first_message_mode": "interruptible",
      "enable_end_call_tool": true,
      "enable_skip_turn_tool": true,
      "enable_play_keypad_touch_tone_tool": true,
      "agent_rules": [
        {
          "title": "value",
          "description": "value",
          "priority": "strict",
          "id": "value",
          "created_at_unix": 123
        }
      ],
      "transfer_rules": [
        {
          "label": "value",
          "phone_number": "value",
          "condition": "value",
          "post_dial_digits": "0",
          "id": "value",
          "created_at_unix": 123
        }
      ],
      "hide_caller_number_on_transfers": true,
      "record_after_transfers": true,
      "procedures": [
        {
          "name": "value",
          "trigger": "value",
          "body": "value",
          "id": "value",
          "created_at_unix": 123,
          "source_key": "value"
        }
      ],
      "enable_secure_mode": true,
      "enable_spam_protection": true
    },
    "deactivated_at_unix": 123,
    "deactivation_reason": "tier_downgrade",
    "name": "value",
    "agent_exists": true,
    "is_transfer_allowed": true,
    "is_active": true
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.ListReceptionistsRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionReceptionistsClient(gotResponse, []byte(`[
  {
    "agent_id": "value",
    "config": {
      "additional_instructions": "value",
      "tone_traits": [
        "value"
      ],
      "details_to_collect": [
        "value"
      ],
      "llm": "gpt-4o-mini",
      "reasoning_effort": "none",
      "disable_rag": true,
      "background_voice_detection": true,
      "turn_eagerness": "patient",
      "voice_id": "value",
      "language": "value",
      "additional_languages": [
        "value"
      ],
      "first_message": "value",
      "first_message_mode": "interruptible",
      "enable_end_call_tool": true,
      "enable_skip_turn_tool": true,
      "enable_play_keypad_touch_tone_tool": true,
      "agent_rules": [
        {
          "title": "value",
          "description": "value",
          "priority": "strict",
          "id": "value",
          "created_at_unix": 123
        }
      ],
      "transfer_rules": [
        {
          "label": "value",
          "phone_number": "value",
          "condition": "value",
          "post_dial_digits": "0",
          "id": "value",
          "created_at_unix": 123
        }
      ],
      "hide_caller_number_on_transfers": true,
      "record_after_transfers": true,
      "procedures": [
        {
          "name": "value",
          "trigger": "value",
          "body": "value",
          "id": "value",
          "created_at_unix": 123,
          "source_key": "value"
        }
      ],
      "enable_secure_mode": true,
      "enable_spam_protection": true
    },
    "deactivated_at_unix": 123,
    "deactivation_reason": "tier_downgrade",
    "name": "value",
    "agent_exists": true,
    "is_transfer_allowed": true,
    "is_active": true
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "agent_id": "value",
    "config": {
      "additional_instructions": "value",
      "tone_traits": [
        "value"
      ],
      "details_to_collect": [
        "value"
      ],
      "llm": "gpt-4o-mini",
      "reasoning_effort": "none",
      "disable_rag": true,
      "background_voice_detection": true,
      "turn_eagerness": "patient",
      "voice_id": "value",
      "language": "value",
      "additional_languages": [
        "value"
      ],
      "first_message": "value",
      "first_message_mode": "interruptible",
      "enable_end_call_tool": true,
      "enable_skip_turn_tool": true,
      "enable_play_keypad_touch_tone_tool": true,
      "agent_rules": [
        {
          "title": "value",
          "description": "value",
          "priority": "strict",
          "id": "value",
          "created_at_unix": 123
        }
      ],
      "transfer_rules": [
        {
          "label": "value",
          "phone_number": "value",
          "condition": "value",
          "post_dial_digits": "0",
          "id": "value",
          "created_at_unix": 123
        }
      ],
      "hide_caller_number_on_transfers": true,
      "record_after_transfers": true,
      "procedures": [
        {
          "name": "value",
          "trigger": "value",
          "body": "value",
          "id": "value",
          "created_at_unix": 123,
          "source_key": "value"
        }
      ],
      "enable_secure_mode": true,
      "enable_spam_protection": true
    },
    "deactivated_at_unix": 123,
    "deactivation_reason": "tier_downgrade",
    "name": "value",
    "agent_exists": true,
    "is_transfer_allowed": true,
    "is_active": true
  }
]`)
	}

	req := receiveWireReceptionReceptionistsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/reception/customer-facing-agents" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/customer-facing-agents")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionReceptionistsClient(t, 422, `{
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

	request := &reception.ListReceptionistsRequest{
		XiWorkspaceId: "value",
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

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionReceptionistsClient(t, 200, `{
  "agent_id": "value",
  "config": {
    "additional_instructions": "value",
    "tone_traits": [
      "value"
    ],
    "details_to_collect": [
      "value"
    ],
    "llm": "gpt-4o-mini",
    "reasoning_effort": "none",
    "disable_rag": true,
    "background_voice_detection": true,
    "turn_eagerness": "patient",
    "voice_id": "value",
    "language": "value",
    "additional_languages": [
      "value"
    ],
    "first_message": "value",
    "first_message_mode": "interruptible",
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "agent_rules": [
      {
        "title": "value",
        "description": "value",
        "priority": "strict",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "transfer_rules": [
      {
        "label": "value",
        "phone_number": "value",
        "condition": "value",
        "post_dial_digits": "0",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "procedures": [
      {
        "name": "value",
        "trigger": "value",
        "body": "value",
        "id": "value",
        "created_at_unix": 123,
        "source_key": "value"
      }
    ],
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "name": "value",
  "agent_exists": true,
  "is_transfer_allowed": true,
  "is_active": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request reception.ReceptionUpdateCustomerFacingAgentRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "agent_id": "value",
  "config": {
    "additional_instructions": "value",
    "tone_traits": [
      "value"
    ],
    "details_to_collect": [
      "value"
    ],
    "llm": "gpt-4o-mini",
    "reasoning_effort": "none",
    "disable_rag": true,
    "background_voice_detection": true,
    "turn_eagerness": "patient",
    "voice_id": "value",
    "language": "value",
    "additional_languages": [
      "value"
    ],
    "first_message": "value",
    "first_message_mode": "interruptible",
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionReceptionistsClient(gotResponse, []byte(`{
  "agent_id": "value",
  "config": {
    "additional_instructions": "value",
    "tone_traits": [
      "value"
    ],
    "details_to_collect": [
      "value"
    ],
    "llm": "gpt-4o-mini",
    "reasoning_effort": "none",
    "disable_rag": true,
    "background_voice_detection": true,
    "turn_eagerness": "patient",
    "voice_id": "value",
    "language": "value",
    "additional_languages": [
      "value"
    ],
    "first_message": "value",
    "first_message_mode": "interruptible",
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "agent_rules": [
      {
        "title": "value",
        "description": "value",
        "priority": "strict",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "transfer_rules": [
      {
        "label": "value",
        "phone_number": "value",
        "condition": "value",
        "post_dial_digits": "0",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "procedures": [
      {
        "name": "value",
        "trigger": "value",
        "body": "value",
        "id": "value",
        "created_at_unix": 123,
        "source_key": "value"
      }
    ],
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "name": "value",
  "agent_exists": true,
  "is_transfer_allowed": true,
  "is_active": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agent_id": "value",
  "config": {
    "additional_instructions": "value",
    "tone_traits": [
      "value"
    ],
    "details_to_collect": [
      "value"
    ],
    "llm": "gpt-4o-mini",
    "reasoning_effort": "none",
    "disable_rag": true,
    "background_voice_detection": true,
    "turn_eagerness": "patient",
    "voice_id": "value",
    "language": "value",
    "additional_languages": [
      "value"
    ],
    "first_message": "value",
    "first_message_mode": "interruptible",
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "agent_rules": [
      {
        "title": "value",
        "description": "value",
        "priority": "strict",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "transfer_rules": [
      {
        "label": "value",
        "phone_number": "value",
        "condition": "value",
        "post_dial_digits": "0",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "procedures": [
      {
        "name": "value",
        "trigger": "value",
        "body": "value",
        "id": "value",
        "created_at_unix": 123,
        "source_key": "value"
      }
    ],
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "name": "value",
  "agent_exists": true,
  "is_transfer_allowed": true,
  "is_active": true
}`)
	}

	req := receiveWireReceptionReceptionistsClient(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/reception/customer-facing-agents/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/customer-facing-agents/value")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsReceptionReceptionistsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionReceptionistsClient(t, 422, `{
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

	var request reception.ReceptionUpdateCustomerFacingAgentRequest
	if err := json.Unmarshal([]byte(`{
  "xi-workspace-id": "value",
  "agent_id": "value",
  "config": {
    "additional_instructions": "value",
    "tone_traits": [
      "value"
    ],
    "details_to_collect": [
      "value"
    ],
    "llm": "gpt-4o-mini",
    "reasoning_effort": "none",
    "disable_rag": true,
    "background_voice_detection": true,
    "turn_eagerness": "patient",
    "voice_id": "value",
    "language": "value",
    "additional_languages": [
      "value"
    ],
    "first_message": "value",
    "first_message_mode": "interruptible",
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "name": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	request.XiWorkspaceId = "value"
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

func TestActivateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireReceptionReceptionistsClient(t, 200, `{
  "agent_id": "value",
  "config": {
    "additional_instructions": "value",
    "tone_traits": [
      "value"
    ],
    "details_to_collect": [
      "value"
    ],
    "llm": "gpt-4o-mini",
    "reasoning_effort": "none",
    "disable_rag": true,
    "background_voice_detection": true,
    "turn_eagerness": "patient",
    "voice_id": "value",
    "language": "value",
    "additional_languages": [
      "value"
    ],
    "first_message": "value",
    "first_message_mode": "interruptible",
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "agent_rules": [
      {
        "title": "value",
        "description": "value",
        "priority": "strict",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "transfer_rules": [
      {
        "label": "value",
        "phone_number": "value",
        "condition": "value",
        "post_dial_digits": "0",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "procedures": [
      {
        "name": "value",
        "trigger": "value",
        "body": "value",
        "id": "value",
        "created_at_unix": 123,
        "source_key": "value"
      }
    ],
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "name": "value",
  "agent_exists": true,
  "is_transfer_allowed": true,
  "is_active": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	request := &reception.ActivateReceptionistsRequest{
		XiWorkspaceId: "value",
	}
	result, err := client.Activate(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualReceptionReceptionistsClient(gotResponse, []byte(`{
  "agent_id": "value",
  "config": {
    "additional_instructions": "value",
    "tone_traits": [
      "value"
    ],
    "details_to_collect": [
      "value"
    ],
    "llm": "gpt-4o-mini",
    "reasoning_effort": "none",
    "disable_rag": true,
    "background_voice_detection": true,
    "turn_eagerness": "patient",
    "voice_id": "value",
    "language": "value",
    "additional_languages": [
      "value"
    ],
    "first_message": "value",
    "first_message_mode": "interruptible",
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "agent_rules": [
      {
        "title": "value",
        "description": "value",
        "priority": "strict",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "transfer_rules": [
      {
        "label": "value",
        "phone_number": "value",
        "condition": "value",
        "post_dial_digits": "0",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "procedures": [
      {
        "name": "value",
        "trigger": "value",
        "body": "value",
        "id": "value",
        "created_at_unix": 123,
        "source_key": "value"
      }
    ],
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "name": "value",
  "agent_exists": true,
  "is_transfer_allowed": true,
  "is_active": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agent_id": "value",
  "config": {
    "additional_instructions": "value",
    "tone_traits": [
      "value"
    ],
    "details_to_collect": [
      "value"
    ],
    "llm": "gpt-4o-mini",
    "reasoning_effort": "none",
    "disable_rag": true,
    "background_voice_detection": true,
    "turn_eagerness": "patient",
    "voice_id": "value",
    "language": "value",
    "additional_languages": [
      "value"
    ],
    "first_message": "value",
    "first_message_mode": "interruptible",
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "agent_rules": [
      {
        "title": "value",
        "description": "value",
        "priority": "strict",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "transfer_rules": [
      {
        "label": "value",
        "phone_number": "value",
        "condition": "value",
        "post_dial_digits": "0",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "procedures": [
      {
        "name": "value",
        "trigger": "value",
        "body": "value",
        "id": "value",
        "created_at_unix": 123,
        "source_key": "value"
      }
    ],
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade",
  "name": "value",
  "agent_exists": true,
  "is_transfer_allowed": true,
  "is_active": true
}`)
	}

	req := receiveWireReceptionReceptionistsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/reception/customer-facing-agents/value/activate" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/reception/customer-facing-agents/value/activate")
	}
	if got := req.Header.Get("xi-workspace-id"); got != "value" {
		t.Fatalf("header %s: got %q want %q", "xi-workspace-id", got, "value")
	}
}

func TestActivateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireReceptionReceptionistsClient(t, 422, `{
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

	request := &reception.ActivateReceptionistsRequest{
		XiWorkspaceId: "value",
	}
	_, err := client.Activate(context.Background(), string("value"), request)
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

type wireRequestReceptionReceptionistsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireReceptionReceptionistsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireReceptionReceptionistsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestReceptionReceptionistsClient) {
	t.Helper()
	recorded := make(chan wireRequestReceptionReceptionistsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestReceptionReceptionistsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireReceptionReceptionistsClient(t *testing.T, recorded <-chan wireRequestReceptionReceptionistsClient) wireRequestReceptionReceptionistsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestReceptionReceptionistsClient{}
	}
}

// wireJSONContainsReceptionReceptionistsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsReceptionReceptionistsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualReceptionReceptionistsClient(got, want)
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

// setWireValueReceptionReceptionistsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueReceptionReceptionistsClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualReceptionReceptionistsClient(got, want []byte) bool {
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
