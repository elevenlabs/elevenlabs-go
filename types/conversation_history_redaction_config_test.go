package types

import (
	"testing"
)

func TestConversationHistoryRedactionConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "entities": [
    "name"
  ],
  "excluded_data_collection_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ConversationHistoryRedactionConfig))
}
