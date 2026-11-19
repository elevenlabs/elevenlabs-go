package types

import (
	"testing"
)

func TestVerifiedVoiceLanguageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "language": "value",
  "model_id": "value",
  "accent": "value",
  "locale": "value",
  "preview_url": "value"
}`
	assertRoundTrip(t, wire, new(VerifiedVoiceLanguageResponseModel))
}

func TestVerifiedVoiceLanguageResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "language": "value",
  "model_id": "value",
  "accent": "value",
  "locale": "value",
  "preview_url": "value"
}`
	assertRoundTrip(t, wire, new(VerifiedVoiceLanguageResponseModel))
}
