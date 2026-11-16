package types

import (
	"testing"
)

func TestScribeUnacceptedTermsErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "unaccepted_terms",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeUnacceptedTermsError))
}
