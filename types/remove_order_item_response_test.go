package types

import (
	"testing"
)

func TestRemoveOrderItemResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "success": true
}`
	assertRoundTrip(t, wire, new(RemoveOrderItemResponse))
}
