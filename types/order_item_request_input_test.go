package types

import (
	"testing"
)

func TestOrderItemRequestInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(OrderItemRequestInput))
}
