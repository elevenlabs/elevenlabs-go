package types

import (
	"testing"
)

func TestExpiresAfterRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "anchor": "created_at",
  "seconds": 123
}`
	assertRoundTrip(t, wire, new(ExpiresAfter))
}
