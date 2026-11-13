package types

import (
	"testing"
)

func TestHistoryItemResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "state": {},
  "voice_category": {}
}`
	assertRoundTrip(t, wire, new(HistoryItemResponse))
}
