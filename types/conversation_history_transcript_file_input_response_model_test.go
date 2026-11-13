package types

import (
	"testing"
)

func TestConversationHistoryTranscriptFileInputResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "file_id": "value",
  "original_filename": "value",
  "mime_type": "value",
  "file_url": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptFileInputResponseModel))
}
