package types

import (
	"testing"
)

func TestOrderResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "order_id": "prodorder_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "name": "value",
  "state": "open",
  "items": [
    {
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
    }
  ],
  "total_amount_usd": 9.99,
  "sandbox": true,
  "created_at": "2024-01-01T00:00:00Z",
  "submitted_at": "2024-01-01T00:00:00Z",
  "paid_at": "2024-01-01T00:00:00Z",
  "accepted_at": "2024-01-01T00:00:00Z",
  "completed_at": "2024-01-01T00:00:00Z",
  "cancel_reason": "value"
}`
	assertRoundTrip(t, wire, new(OrderResponse))
}
