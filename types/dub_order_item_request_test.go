package types

import (
	"testing"
)

func TestDubOrderItemRequestRoundTrip(t *testing.T) {
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
	assertRoundTrip(t, wire, new(DubOrderItemRequest))
}

func TestDubOrderItemRequestRoundTrip2(t *testing.T) {
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
	assertRoundTrip(t, wire, new(DubOrderItemRequest))
}
