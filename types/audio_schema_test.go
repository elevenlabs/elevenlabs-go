package types

import (
	"testing"
)

func TestAudioSchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "audio"
}`
	assertRoundTrip(t, wire, new(AudioSchema))
}
