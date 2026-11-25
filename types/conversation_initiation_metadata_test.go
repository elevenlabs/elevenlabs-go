package types

import (
	"testing"
)

func TestConversationInitiationMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_initiation_metadata_event": {
    "conversation_id": "value",
    "agent_output_audio_format": "pcm_8000",
    "user_input_audio_format": "pcm_8000"
  },
  "type": "conversation_initiation_metadata"
}`
	assertRoundTrip(t, wire, new(ConversationInitiationMetadata))
}
