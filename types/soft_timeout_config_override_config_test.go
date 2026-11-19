package types

import (
	"testing"
)

func TestSoftTimeoutConfigOverrideConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message": true,
  "additional_soft_timeout_messages": true
}`
	assertRoundTrip(t, wire, new(SoftTimeoutConfigOverrideConfig))
}
