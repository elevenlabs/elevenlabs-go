package types

import (
	"testing"
)

func TestCommittedTranscriptEntitiesRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "committed_transcript_entities",
  "text": "value",
  "entities": [
    {
      "text": "value",
      "entity_type": "value",
      "start_char": 123,
      "end_char": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(CommittedTranscriptEntities))
}
