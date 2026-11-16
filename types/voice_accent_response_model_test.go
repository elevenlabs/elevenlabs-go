package types

import (
	"testing"
)

func TestVoiceAccentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "accent": "value",
  "language": "value",
  "code": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(VoiceAccentResponseModel))
}
