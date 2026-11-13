package types

import (
	"testing"
)

func TestAutoGenerationMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "stable_topic_id": "value",
  "criteria_id": "value",
  "topic_label": "value",
  "source_conversation_id": "value",
  "generated_at_unix": 123
}`
	assertRoundTrip(t, wire, new(AutoGenerationMetadata))
}
