package types

import (
	"testing"
)

func TestSubtitlingPaymentIssueErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "payment_issue",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(SubtitlingPaymentIssueError))
}
