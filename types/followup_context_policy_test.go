package types

import (
	"testing"
)

func TestFollowupContextPolicyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "call_type": true,
  "previous_conversation_id": true,
  "previous_summary": true,
  "previous_transcript": true
}`
	assertRoundTrip(t, wire, new(FollowupContextPolicy))
}

func TestFollowupContextPolicyRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "call_type": true,
  "previous_conversation_id": true,
  "previous_summary": true,
  "previous_transcript": true
}`
	assertRoundTrip(t, wire, new(FollowupContextPolicy))
}
