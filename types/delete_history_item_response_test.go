package types

import (
	"testing"
)

func TestDeleteHistoryItemResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteHistoryItemResponse))
}
