package types

import (
	"testing"
)

func TestCommittedTranscriptEntitiesEntitiesItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "entity_type": "value",
  "start_char": 123,
  "end_char": 123
}`
	assertRoundTrip(t, wire, new(CommittedTranscriptEntitiesEntitiesItem))
}
