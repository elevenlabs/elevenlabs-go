package types

import (
	"testing"
)

func TestConversationDeletionSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "deletion_time_unix_secs": 123,
  "deleted_logs_at_time_unix_secs": 123,
  "deleted_audio_at_time_unix_secs": 123,
  "deleted_transcript_at_time_unix_secs": 123,
  "delete_transcript_and_pii": true,
  "delete_audio": true
}`
	assertRoundTrip(t, wire, new(ConversationDeletionSettings))
}
