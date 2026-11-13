package types

import (
	"testing"
)

func TestEditVoiceSettingsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(EditVoiceSettingsResponseModel))
}
