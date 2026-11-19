package types

import (
	"testing"
)

func TestOTelScopeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "version": "value"
}`
	assertRoundTrip(t, wire, new(OTelScope))
}
