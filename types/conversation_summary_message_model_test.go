package types

import (
	"testing"
)

func TestConversationSummaryMessageModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "role": "user",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(ConversationSummaryMessageModel))
}
