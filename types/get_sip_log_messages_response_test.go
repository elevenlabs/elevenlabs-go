package types

import (
	"testing"
)

func TestGetSipLogMessagesResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(GetSipLogMessagesResponse))
}

func TestGetSipLogMessagesResponseRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(GetSipLogMessagesResponse))
}
