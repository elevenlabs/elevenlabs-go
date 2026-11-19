package types

import (
	"testing"
)

func TestUpsertOrderItemRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
  "item_id": "proditem_aaaaaaaaaaaaaaaaaaaaaaaaaa"
}`
	assertRoundTrip(t, wire, new(UpsertOrderItemRequest))
}
