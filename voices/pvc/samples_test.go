package pvc

import (
	"testing"
)

func TestBodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "remove_background_noise": true,
  "selected_speaker_ids": [
    "value"
  ],
  "trim_start_time": 123,
  "trim_end_time": 123,
  "file_name": "value"
}`
	assertRoundTrip(t, wire, new(BodyUpdatePvcVoiceSampleV1VoicesPvcVoiceIdSamplesSampleIdPost))
}
