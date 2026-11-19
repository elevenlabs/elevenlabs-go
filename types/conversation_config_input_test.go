package types

import (
	"testing"
)

func TestConversationConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text_only": true,
  "max_duration_seconds": 123,
  "client_events": [
    "conversation_initiation_metadata"
  ],
  "file_input": {
    "enabled": true,
    "max_files_in_memory": 30,
    "max_files_per_conversation": 123
  },
  "monitoring_enabled": true,
  "monitoring_events": [
    "conversation_initiation_metadata"
  ],
  "dtmf_input_settings": {
    "dtmf_input_timeout": 9.99,
    "hash_terminator": true,
    "redact_input": true
  },
  "background_sound": {
    "source_type": "preset",
    "source_id": "office2",
    "volume": 0.5,
    "crossfade_loop": true
  },
  "source_attribution": true
}`
	assertRoundTrip(t, wire, new(ConversationConfigInput))
}

func TestConversationConfigInputRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text_only": true,
  "max_duration_seconds": 123,
  "client_events": [
    "conversation_initiation_metadata"
  ],
  "file_input": {
    "enabled": true,
    "max_files_in_memory": 30,
    "max_files_per_conversation": 123
  },
  "monitoring_enabled": true,
  "monitoring_events": [
    "conversation_initiation_metadata"
  ],
  "dtmf_input_settings": {
    "dtmf_input_timeout": 9.99,
    "hash_terminator": true,
    "redact_input": true
  },
  "background_sound": {
    "source_type": "preset",
    "source_id": "office2",
    "volume": 0.5,
    "crossfade_loop": true
  },
  "source_attribution": true
}`
	assertRoundTrip(t, wire, new(ConversationConfigInput))
}
