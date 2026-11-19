package types

import (
	"testing"
)

func TestTicketCommentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "comment": "value",
  "created_at_unix_secs": 123,
  "owner_user_id": "value"
}`
	assertRoundTrip(t, wire, new(TicketCommentResponseModel))
}
