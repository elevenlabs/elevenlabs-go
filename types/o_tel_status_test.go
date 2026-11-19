package types

import (
	"testing"
)

func TestOTelStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "code": 123,
  "message": "value"
}`
	assertRoundTrip(t, wire, new(OTelStatus))
}
