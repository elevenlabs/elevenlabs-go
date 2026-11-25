package types

import (
	"testing"
)

func TestConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "pronunciation_dictionary_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem))
}
