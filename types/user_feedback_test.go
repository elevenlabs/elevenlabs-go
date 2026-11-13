package types

import (
	"testing"
)

func TestUserFeedbackRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "score": "like",
  "time_in_call_secs": 123
}`
	assertRoundTrip(t, wire, new(UserFeedback))
}

func TestUserFeedbackRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "score": "like",
  "time_in_call_secs": 123
}`
	assertRoundTrip(t, wire, new(UserFeedback))
}

func TestUserFeedbackRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "score": "like",
  "time_in_call_secs": 123
}`
	assertRoundTrip(t, wire, new(UserFeedback))
}
