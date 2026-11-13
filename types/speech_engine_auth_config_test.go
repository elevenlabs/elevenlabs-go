package types

import (
	"testing"
)

func TestSpeechEngineAuthConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enable_auth": true
}`
	assertRoundTrip(t, wire, new(SpeechEngineAuthConfig))
}
