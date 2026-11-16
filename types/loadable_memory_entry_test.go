package types

import (
	"testing"
)

func TestLoadableMemoryEntryRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "entry_id": "value",
  "summary": "value",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(LoadableMemoryEntry))
}
