package types

import (
	"testing"
)

func TestStudioAgentSettingsModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_settings": {
    "key": {
      "skip_confirmation": true
    }
  }
}`
	assertRoundTrip(t, wire, new(StudioAgentSettingsModel))
}

func TestStudioAgentSettingsModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_settings": {
    "key": {
      "skip_confirmation": true
    }
  }
}`
	assertRoundTrip(t, wire, new(StudioAgentSettingsModel))
}
