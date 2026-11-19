package types

import (
	"testing"
)

func TestConversationHistoryTranscriptSystemToolResultCommonModelInputResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "end_call_success",
  "status": "success",
  "reason": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptSystemToolResultCommonModelInputResult))
}
