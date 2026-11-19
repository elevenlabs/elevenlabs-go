package types

import (
	"testing"
)

func TestConvAiWebhooksRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "post_call_webhook_id": "value",
  "events": [
    "transcript"
  ],
  "transcript_format": "json",
  "exclude_transcript": true,
  "send_audio": true
}`
	assertRoundTrip(t, wire, new(ConvAiWebhooks))
}

func TestConvAiWebhooksRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "post_call_webhook_id": "value",
  "events": [
    "transcript"
  ],
  "transcript_format": "json",
  "exclude_transcript": true,
  "send_audio": true
}`
	assertRoundTrip(t, wire, new(ConvAiWebhooks))
}
