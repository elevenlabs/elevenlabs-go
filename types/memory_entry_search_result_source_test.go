package types

import (
	"testing"
)

func TestMemoryEntrySearchResultSourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "conversation",
  "conversation_id": "value"
}`
	assertRoundTrip(t, wire, new(MemoryEntrySearchResultSource))
}
