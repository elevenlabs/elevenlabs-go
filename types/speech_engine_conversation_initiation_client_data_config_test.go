package types

import (
	"testing"
)

func TestSpeechEngineConversationInitiationClientDataConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "first_message": true
}`
	assertRoundTrip(t, wire, new(SpeechEngineConversationInitiationClientDataConfig))
}
