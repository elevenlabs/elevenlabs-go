package types

import (
	"testing"
)

func TestVoiceStatisticsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "project_voice_ref_id": "value",
  "characters_unconverted": 123,
  "characters_converted": 123,
  "credits_needed_to_convert": 123,
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(VoiceStatisticsResponseModel))
}
