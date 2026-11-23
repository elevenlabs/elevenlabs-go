package types

import (
	"testing"
)

func TestFeedbackRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "feedback",
  "event_id": 123,
  "score": "like"
}`
	assertRoundTrip(t, wire, new(Feedback))
}
