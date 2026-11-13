package types

import (
	"testing"
)

func TestRagDocumentIndexUsageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "used_bytes": 123
}`
	assertRoundTrip(t, wire, new(RagDocumentIndexUsage))
}
