package types

import (
	"testing"
)

func TestReceptionServiceScrapableRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pricing_mode": "fixed",
  "groups": [
    {
      "label": "value",
      "rank": 123
    }
  ],
  "ungrouped_rank": 123,
  "name": "value",
  "description": "value",
  "variants": [
    {
      "id": "value",
      "duration_minutes": 123,
      "price": 9.99,
      "is_active": true
    }
  ],
  "kind": "classic",
  "buffer_minutes": 123,
  "service_location": "business",
  "add_ons": [
    {
      "id": "value",
      "name": "value",
      "description": "value",
      "add_duration_minutes": 123,
      "add_price": 9.99
    }
  ],
  "max_participants": 123
}`
	assertRoundTrip(t, wire, new(ReceptionServiceScrapable))
}
