package types

import (
	"testing"
)

func TestTicketPriorityChangeResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "priority": "low",
  "changed_by_user_id": "value",
  "changed_at_unix_secs": 123
}`
	assertRoundTrip(t, wire, new(TicketPriorityChangeResponseModel))
}
