package types

import (
	"testing"
)

func TestTurnCommentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "turn_index": 123,
  "comment": "value",
  "created_at_unix_secs": 123,
  "owner_user_id": "value"
}`
	assertRoundTrip(t, wire, new(TurnCommentResponseModel))
}
