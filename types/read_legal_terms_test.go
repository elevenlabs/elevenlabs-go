package types

import (
	"testing"
)

func TestReadLegalTermsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "terms": "value",
  "start_date": "value",
  "end_date": "value"
}`
	assertRoundTrip(t, wire, new(ReadLegalTerms))
}
