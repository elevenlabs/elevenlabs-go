package types

import (
	"testing"
)

func TestConversationHistoryErrorCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "code": 123,
  "reason": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryErrorCommonModel))
}
