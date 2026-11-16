package types

import (
	"testing"
)

func TestTurnCommentRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "turn_index": 123,
  "comment": "value"
}`
	assertRoundTrip(t, wire, new(TurnCommentRequestModel))
}
