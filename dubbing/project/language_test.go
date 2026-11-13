package project

import (
	"testing"
)

func TestBodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "target_language": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "translations": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(BodyCreateDubbingLanguageTargetV1DubbingProjectProjectIdLanguagePost))
}
