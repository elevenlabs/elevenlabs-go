package types

import (
	"testing"
)

func TestReceptionProductRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`
	assertRoundTrip(t, wire, new(ReceptionProduct))
}

func TestReceptionProductRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pricing_mode": "fixed",
  "intake_questions": [
    {
      "id": "value",
      "label": "value",
      "kind": "text",
      "options": [
        "value"
      ],
      "required": true,
      "is_active": true
    }
  ],
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "price": 9.99,
  "is_active": true,
  "orderable": true,
  "allowed_fulfillment_types": [
    "pickup"
  ],
  "location_ids": [
    "value"
  ],
  "stock": {
    "key": 123
  },
  "id": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123
}`
	assertRoundTrip(t, wire, new(ReceptionProduct))
}
