package types

import (
	"testing"
)

func TestTranslateUnacceptedTermsErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "unaccepted_terms",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(TranslateUnacceptedTermsError))
}
