package types

import (
	"testing"
)

func TestVoicemailDetectionToolConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "voicemail_detection",
  "voicemail_message": "value"
}`
	assertRoundTrip(t, wire, new(VoicemailDetectionToolConfig))
}
