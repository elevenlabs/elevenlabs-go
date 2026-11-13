package types

import (
	"testing"
)

func TestBodyCreateANewVoiceFromVoicePreviewV1TextToVoicePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_name": "value",
  "voice_description": "valuexxxxxxxxxxxxxxx",
  "generated_voice_id": "value",
  "labels": {
    "key": "value"
  },
  "played_not_selected_voice_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(BodyCreateANewVoiceFromVoicePreviewV1TextToVoicePost))
}
