package types

import (
	"testing"
)

func TestOTelAttributeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "key": "value",
  "value": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(OTelAttribute))
}
