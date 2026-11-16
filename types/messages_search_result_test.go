package types

import (
	"testing"
)

func TestMessagesSearchResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_id": "value",
  "agent_id": "value",
  "agent_name": "value",
  "transcript_index": 123,
  "chunk_text": "value",
  "chunk_highlights": [
    {
      "value": "value",
      "is_hit": true
    }
  ],
  "score": 9.99,
  "conversation_start_time_unix_secs": 123
}`
	assertRoundTrip(t, wire, new(MessagesSearchResult))
}
