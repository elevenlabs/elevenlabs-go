package types

import (
	"testing"
)

func TestStudioAgentToolSettingsModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "skip_confirmation": true
}`
	assertRoundTrip(t, wire, new(StudioAgentToolSettingsModel))
}
