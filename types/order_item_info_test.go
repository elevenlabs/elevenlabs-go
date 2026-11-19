package types

import (
	"testing"
)

func TestOrderItemInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "item_id": "proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "item": {
    "kind": "dub",
    "media_id": "prodmedia_aaaaaaaaaaaaaaaaaaaaaaaaaa",
    "source_language": "value",
    "destination_languages": [
      "value"
    ],
    "include_captions": true,
    "include_source_captions": true,
    "instructions": "value",
    "captions_sdh": true
  },
  "quote": {
    "amount_usd": 9.99
  }
}`
	assertRoundTrip(t, wire, new(OrderItemInfo))
}
