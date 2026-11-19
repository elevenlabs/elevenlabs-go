package types

import (
	"testing"
)

func TestSipLogMessageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(SipLogMessage))
}
