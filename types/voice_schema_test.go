package types

import (
	"testing"
)

func TestVoiceSchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "voice"
}`
	assertRoundTrip(t, wire, new(VoiceSchema))
}
