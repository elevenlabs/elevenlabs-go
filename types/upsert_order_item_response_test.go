package types

import (
	"testing"
)

func TestUpsertOrderItemResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "item_id": "proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "quote": {
    "amount_usd": 9.99
  }
}`
	assertRoundTrip(t, wire, new(UpsertOrderItemResponse))
}
