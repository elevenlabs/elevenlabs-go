package types

import (
	"testing"
)

func TestSpeechToTextWebhookResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message": "value",
  "request_id": "value",
  "transcription_id": "value"
}`
	assertRoundTrip(t, wire, new(SpeechToTextWebhookResponseModel))
}
