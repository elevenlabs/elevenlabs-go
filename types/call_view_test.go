package types

import (
	"testing"
)

func TestCallViewRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "intent_id": "value",
  "lead_id": "value",
  "intent_type": "initial",
  "cold_attempt": 123,
  "status": "queued",
  "phone_number": "value",
  "scheduled_time_unix": 123,
  "conversation_id": "value",
  "outcome": "pending",
  "call_result": "answered",
  "outcome_at_unix": 123,
  "decision_disposition": "completed",
  "collected_dynamic_variables": {
    "key": {}
  },
  "delivery": {
    "status": "pending",
    "stage": "queued",
    "sent_at_unix": 123,
    "delivered_at_unix": 123,
    "read_at_unix": 123,
    "replied_at_unix": 123,
    "failed_at_unix": 123,
    "error_message": "value"
  }
}`
	assertRoundTrip(t, wire, new(CallView))
}
