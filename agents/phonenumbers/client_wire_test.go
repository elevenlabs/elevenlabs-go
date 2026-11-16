package phonenumbers

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

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsPhonenumbers(t, 200, `[
  {
    "phone_number": "value",
    "label": "value",
    "supports_inbound": true,
    "supports_outbound": true,
    "phone_number_id": "value",
    "assigned_agent": {
      "agent_id": "value",
      "agent_name": "value",
      "environment": "value",
      "branch_id": "value"
    },
    "provider": "twilio"
  }
]`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	agentIdValue := string("value")
	branchIdValue := string("value")
	request := &agents.ListPhoneNumbersRequest{
		AgentId:  &agentIdValue,
		BranchId: &branchIdValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsPhonenumbers(gotResponse, []byte(`[
  {
    "phone_number": "value",
    "label": "value",
    "supports_inbound": true,
    "supports_outbound": true,
    "phone_number_id": "value",
    "assigned_agent": {
      "agent_id": "value",
      "agent_name": "value",
      "environment": "value",
      "branch_id": "value"
    },
    "provider": "twilio"
  }
]`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `[
  {
    "phone_number": "value",
    "label": "value",
    "supports_inbound": true,
    "supports_outbound": true,
    "phone_number_id": "value",
    "assigned_agent": {
      "agent_id": "value",
      "agent_name": "value",
      "environment": "value",
      "branch_id": "value"
    },
    "provider": "twilio"
  }
]`)
	}

	req := receiveWireAgentsPhonenumbers(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/phone-numbers" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/phone-numbers")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "branch_id", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsPhonenumbers(t, 422, `{
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
	branchIdValue := string("value")
	request := &agents.ListPhoneNumbersRequest{
		AgentId:  &agentIdValue,
		BranchId: &branchIdValue,
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

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsPhonenumbers(t, 200, `{
  "phone_number_id": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.CreatePhoneNumbersRequestBody
	if err := json.Unmarshal([]byte(`{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "provider": "twilio",
  "agent_id": "value",
  "sid": "value",
  "token": "value",
  "account_auth_token": "value",
  "region_config": {
    "region_id": "us1",
    "token": "value",
    "edge_location": "ashburn"
  },
  "enable_sms": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsPhonenumbers(gotResponse, []byte(`{
  "phone_number_id": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "phone_number_id": "value"
}`)
	}

	req := receiveWireAgentsPhonenumbers(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/phone-numbers" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/phone-numbers")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsPhonenumbers(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsPhonenumbers(t, 422, `{
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

	var request agents.CreatePhoneNumbersRequestBody
	if err := json.Unmarshal([]byte(`{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "provider": "twilio",
  "agent_id": "value",
  "sid": "value",
  "token": "value",
  "account_auth_token": "value",
  "region_config": {
    "region_id": "us1",
    "token": "value",
    "edge_location": "ashburn"
  },
  "enable_sms": true
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	_, err := client.Create(context.Background(), &request)
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

	server, recorded := serveWireAgentsPhonenumbers(t, 200, `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "twilio"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsPhonenumbers(gotResponse, []byte(`{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "twilio"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "twilio"
}`)
	}

	req := receiveWireAgentsPhonenumbers(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/phone-numbers/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/phone-numbers/value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsPhonenumbers(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"))
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

	server, recorded := serveWireAgentsPhonenumbers(t, 200, `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "twilio"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request agents.UpdatePhoneNumberRequest
	if err := json.Unmarshal([]byte(`{
  "phone_number_id": "value",
  "agent_id": "value",
  "label": "value",
  "inbound_trunk_config": {
    "allowed_addresses": [
      "value"
    ],
    "allowed_numbers": [
      "value"
    ],
    "media_encryption": "disabled",
    "credentials": {
      "username": "value",
      "password": "value"
    },
    "remote_domains": [
      "value"
    ],
    "attributes_to_headers": {
      "key": "value"
    }
  },
  "outbound_trunk_config": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "credentials": {
      "username": "value",
      "password": "value"
    },
    "enabled_codecs": [
      "G722/8000"
    ]
  },
  "livekit_stack": "standard",
  "store_sip_messages": true,
  "environment": "value",
  "branch_id": "value"
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
	if !wireJSONBytesEqualAgentsPhonenumbers(gotResponse, []byte(`{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "twilio"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "twilio"
}`)
	}

	req := receiveWireAgentsPhonenumbers(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/phone-numbers/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/phone-numbers/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsPhonenumbers(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsPhonenumbers(t, 422, `{
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

	var request agents.UpdatePhoneNumberRequest
	if err := json.Unmarshal([]byte(`{
  "phone_number_id": "value",
  "agent_id": "value",
  "label": "value",
  "inbound_trunk_config": {
    "allowed_addresses": [
      "value"
    ],
    "allowed_numbers": [
      "value"
    ],
    "media_encryption": "disabled",
    "credentials": {
      "username": "value",
      "password": "value"
    },
    "remote_domains": [
      "value"
    ],
    "attributes_to_headers": {
      "key": "value"
    }
  },
  "outbound_trunk_config": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "credentials": {
      "username": "value",
      "password": "value"
    },
    "enabled_codecs": [
      "G722/8000"
    ]
  },
  "livekit_stack": "standard",
  "store_sip_messages": true,
  "environment": "value",
  "branch_id": "value"
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

	server, recorded := serveWireAgentsPhonenumbers(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsPhonenumbers(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsPhonenumbers(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/phone-numbers/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/phone-numbers/value")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsPhonenumbers(t, 422, `{
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

	_, err := client.Delete(context.Background(), string("value"))
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

func TestListV2Wire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsPhonenumbers(t, 200, `{
  "phone_numbers": [
    {
      "phone_number": "value",
      "label": "value",
      "supports_inbound": true,
      "supports_outbound": true,
      "phone_number_id": "value",
      "assigned_agent": {
        "agent_id": "value",
        "agent_name": "value",
        "environment": "value",
        "branch_id": "value"
      },
      "provider": "twilio"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(123)
	searchValue := string("value")
	labelValue := string("value")
	phoneNumberValue := string("value")
	supportsOutboundValue := bool(true)
	agentIdValue := string("value")
	branchIdValue := string("value")
	cursorValue := string("value")
	request := &agents.ListV2PhoneNumbersRequest{
		PageSize:         &pageSizeValue,
		Search:           &searchValue,
		Label:            &labelValue,
		PhoneNumber:      &phoneNumberValue,
		SupportsOutbound: &supportsOutboundValue,
		AgentId:          &agentIdValue,
		BranchId:         &branchIdValue,
		Cursor:           &cursorValue,
	}
	result, err := client.ListV2(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsPhonenumbers(gotResponse, []byte(`{
  "phone_numbers": [
    {
      "phone_number": "value",
      "label": "value",
      "supports_inbound": true,
      "supports_outbound": true,
      "phone_number_id": "value",
      "assigned_agent": {
        "agent_id": "value",
        "agent_name": "value",
        "environment": "value",
        "branch_id": "value"
      },
      "provider": "twilio"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "phone_numbers": [
    {
      "phone_number": "value",
      "label": "value",
      "supports_inbound": true,
      "supports_outbound": true,
      "phone_number_id": "value",
      "assigned_agent": {
        "agent_id": "value",
        "agent_name": "value",
        "environment": "value",
        "branch_id": "value"
      },
      "provider": "twilio"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsPhonenumbers(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/v2/phone-numbers" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/v2/phone-numbers")
	}
	if got := req.Query.Get("page_size"); got != "123" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "123")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("label"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "label", got, "value")
	}
	if got := req.Query.Get("phone_number"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "phone_number", got, "value")
	}
	if got := req.Query.Get("supports_outbound"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "supports_outbound", got, "true")
	}
	if got := req.Query.Get("agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "agent_id", got, "value")
	}
	if got := req.Query.Get("branch_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "branch_id", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListV2WireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsPhonenumbers(t, 422, `{
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

	pageSizeValue := int(123)
	searchValue := string("value")
	labelValue := string("value")
	phoneNumberValue := string("value")
	supportsOutboundValue := bool(true)
	agentIdValue := string("value")
	branchIdValue := string("value")
	cursorValue := string("value")
	request := &agents.ListV2PhoneNumbersRequest{
		PageSize:         &pageSizeValue,
		Search:           &searchValue,
		Label:            &labelValue,
		PhoneNumber:      &phoneNumberValue,
		SupportsOutbound: &supportsOutboundValue,
		AgentId:          &agentIdValue,
		BranchId:         &branchIdValue,
		Cursor:           &cursorValue,
	}
	_, err := client.ListV2(context.Background(), request)
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

func TestGetSipMessagesWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsPhonenumbers(t, 200, `{
  "sip_messages": [
    {
      "call_id": "value",
      "phone_numbers": [
        "value"
      ],
      "local_address": "value",
      "remote_address": "value",
      "transport": "value",
      "raw_message": "value",
      "error_message": "value",
      "direction": "in",
      "created_at_unix_micro": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	pageSizeValue := int(20)
	cursorValue := string("value")
	request := &agents.GetSipMessagesPhoneNumbersRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	result, err := client.GetSipMessages(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsPhonenumbers(gotResponse, []byte(`{
  "sip_messages": [
    {
      "call_id": "value",
      "phone_numbers": [
        "value"
      ],
      "local_address": "value",
      "remote_address": "value",
      "transport": "value",
      "raw_message": "value",
      "error_message": "value",
      "direction": "in",
      "created_at_unix_micro": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "sip_messages": [
    {
      "call_id": "value",
      "phone_numbers": [
        "value"
      ],
      "local_address": "value",
      "remote_address": "value",
      "transport": "value",
      "raw_message": "value",
      "error_message": "value",
      "direction": "in",
      "created_at_unix_micro": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsPhonenumbers(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/phone-numbers/value/sip-messages" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/phone-numbers/value/sip-messages")
	}
	if got := req.Query.Get("page_size"); got != "20" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "20")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestGetSipMessagesWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsPhonenumbers(t, 422, `{
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

	pageSizeValue := int(20)
	cursorValue := string("value")
	request := &agents.GetSipMessagesPhoneNumbersRequest{
		PageSize: &pageSizeValue,
		Cursor:   &cursorValue,
	}
	_, err := client.GetSipMessages(context.Background(), string("value"), request)
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

type wireRequestAgentsPhonenumbers struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsPhonenumbers records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsPhonenumbers(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsPhonenumbers) {
	t.Helper()
	recorded := make(chan wireRequestAgentsPhonenumbers, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsPhonenumbers{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireAgentsPhonenumbers(t *testing.T, recorded <-chan wireRequestAgentsPhonenumbers) wireRequestAgentsPhonenumbers {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsPhonenumbers{}
	}
}

// wireJSONContainsAgentsPhonenumbers reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsPhonenumbers(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsPhonenumbers(got, want)
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

// setWireValueAgentsPhonenumbers sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsPhonenumbers(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualAgentsPhonenumbers(got, want []byte) bool {
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
