package types

import (
	"testing"
)

func TestReceptionAddOnRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "description": "value",
  "add_duration_minutes": 123,
  "add_price": 9.99
}`
	assertRoundTrip(t, wire, new(ReceptionAddOn))
}

func TestReceptionAddOnRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "description": "value",
  "add_duration_minutes": 123,
  "add_price": 9.99
}`
	assertRoundTrip(t, wire, new(ReceptionAddOn))
}
