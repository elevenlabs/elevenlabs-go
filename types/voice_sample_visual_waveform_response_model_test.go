package types

import (
	"testing"
)

func TestVoiceSampleVisualWaveformResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "sample_id": "value",
  "visual_waveform": [
    9.99
  ]
}`
	assertRoundTrip(t, wire, new(VoiceSampleVisualWaveformResponseModel))
}
