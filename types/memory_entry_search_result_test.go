package types

import (
	"testing"
)

func TestMemoryEntrySearchResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "entry_id": "value",
  "version": 123,
  "summary": "value",
  "text": "value",
  "source": {
    "type": "conversation",
    "conversation_id": "value"
  }
}`
	assertRoundTrip(t, wire, new(MemoryEntrySearchResult))
}
