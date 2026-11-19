package types

import (
	"testing"
)

func TestTicketStatusChangeResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "open",
  "changed_by_user_id": "value",
  "changed_at_unix_secs": 123
}`
	assertRoundTrip(t, wire, new(TicketStatusChangeResponseModel))
}
