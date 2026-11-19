package types

import (
	"testing"
)

func TestSoftTimeoutConfigOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message": "value",
  "additional_soft_timeout_messages": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(SoftTimeoutConfigOverride))
}
