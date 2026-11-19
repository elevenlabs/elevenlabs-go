package types

import (
	"testing"
)

func TestDateModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "year": 2026,
  "month": 12,
  "day": 31
}`
	assertRoundTrip(t, wire, new(DateModel))
}

func TestDateModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "year": 2026,
  "month": 12,
  "day": 31
}`
	assertRoundTrip(t, wire, new(DateModel))
}
