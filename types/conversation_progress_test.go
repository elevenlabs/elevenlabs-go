package types

import (
	"testing"
)

func TestConversationProgressRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "scheduled": 123,
  "scheduled_by_type": {
    "key": 123
  },
  "awaiting_permission": 123,
  "in_progress": 123,
  "awaiting_analysis": 123
}`
	assertRoundTrip(t, wire, new(ConversationProgress))
}
