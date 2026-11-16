package types

import (
	"testing"
)

func TestConversationHistoryRagUsageCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "usage_count": 123,
  "embedding_model": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryRagUsageCommonModel))
}
