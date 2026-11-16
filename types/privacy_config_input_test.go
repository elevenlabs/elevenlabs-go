package types

import (
	"testing"
)

func TestPrivacyConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "record_voice": true,
  "retention_days": 123,
  "delete_transcript_and_pii": true,
  "delete_audio": true,
  "apply_to_existing_conversations": true,
  "zero_retention_mode": true,
  "conversation_history_redaction": {
    "enabled": true,
    "entities": [
      "name"
    ],
    "excluded_data_collection_ids": [
      "value"
    ]
  }
}`
	assertRoundTrip(t, wire, new(PrivacyConfigInput))
}

func TestPrivacyConfigInputRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "record_voice": true,
  "retention_days": 123,
  "delete_transcript_and_pii": true,
  "delete_audio": true,
  "apply_to_existing_conversations": true,
  "zero_retention_mode": true,
  "conversation_history_redaction": {
    "enabled": true,
    "entities": [
      "name"
    ],
    "excluded_data_collection_ids": [
      "value"
    ]
  }
}`
	assertRoundTrip(t, wire, new(PrivacyConfigInput))
}
