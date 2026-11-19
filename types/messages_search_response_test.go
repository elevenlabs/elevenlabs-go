package types

import (
	"testing"
)

func TestMessagesSearchResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "meta": {
    "total": 123,
    "page": 123,
    "page_size": 123
  },
  "results": [
    {
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
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(MessagesSearchResponse))
}
