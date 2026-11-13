package types

import (
	"testing"
)

func TestVoiceMailDetectionResultSuccessModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "voicemail_detection_success",
  "status": "success",
  "voicemail_message": "value",
  "reason": "value"
}`
	assertRoundTrip(t, wire, new(VoiceMailDetectionResultSuccessModel))
}

func TestVoiceMailDetectionResultSuccessModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "voicemail_detection_success",
  "status": "success",
  "voicemail_message": "value",
  "reason": "value"
}`
	assertRoundTrip(t, wire, new(VoiceMailDetectionResultSuccessModel))
}
