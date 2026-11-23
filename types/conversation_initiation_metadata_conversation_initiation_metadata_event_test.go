package types

import (
	"testing"
)

func TestConversationInitiationMetadataConversationInitiationMetadataEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_id": "value",
  "agent_output_audio_format": "pcm_8000",
  "user_input_audio_format": "pcm_8000"
}`
	assertRoundTrip(t, wire, new(ConversationInitiationMetadataConversationInitiationMetadataEvent))
}
