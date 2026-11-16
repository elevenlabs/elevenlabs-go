package types

import (
	"testing"
)

func TestTranslateSessionStartedRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "session_started",
  "session_id": "value",
  "client_session_id": "value"
}`
	assertRoundTrip(t, wire, new(TranslateSessionStarted))
}
