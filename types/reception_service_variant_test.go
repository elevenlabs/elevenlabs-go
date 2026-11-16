package types

import (
	"testing"
)

func TestReceptionServiceVariantRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "duration_minutes": 123,
  "price": 9.99,
  "is_active": true
}`
	assertRoundTrip(t, wire, new(ReceptionServiceVariant))
}

func TestReceptionServiceVariantRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "duration_minutes": 123,
  "price": 9.99,
  "is_active": true
}`
	assertRoundTrip(t, wire, new(ReceptionServiceVariant))
}
