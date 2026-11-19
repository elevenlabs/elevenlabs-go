package types

import (
	"testing"
)

func TestReceptionServicePromotionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "target": "service",
  "variant_id": "value",
  "add_on_id": "value",
  "promotional_price": 9.99,
  "days_of_week": [
    6
  ],
  "start_date": "2024-01-01",
  "end_date": "2024-01-01",
  "is_active": true
}`
	assertRoundTrip(t, wire, new(ReceptionServicePromotion))
}
