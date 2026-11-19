package types

import (
	"testing"
)

func TestFlattenedConfigOverridesRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "language": "value",
  "first_message": "value",
  "prompt": "value",
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(FlattenedConfigOverrides))
}
