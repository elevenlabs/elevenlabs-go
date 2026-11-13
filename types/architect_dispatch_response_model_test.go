package types

import (
	"testing"
)

func TestArchitectDispatchResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "architect_chat_id": "value",
  "user_id": "value",
  "dispatched_at_unix_secs": 123
}`
	assertRoundTrip(t, wire, new(ArchitectDispatchResponseModel))
}
