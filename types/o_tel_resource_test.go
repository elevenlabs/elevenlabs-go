package types

import (
	"testing"
)

func TestOTelResourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "attributes": [
    {
      "key": "value",
      "value": {
        "key": "value"
      }
    }
  ]
}`
	assertRoundTrip(t, wire, new(OTelResource))
}
