package types

import (
	"testing"
)

func TestDubbingLanguageResponseVoiceSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "cloning_strength": 10
}`
	assertRoundTrip(t, wire, new(DubbingLanguageResponseVoiceSettings))
}
