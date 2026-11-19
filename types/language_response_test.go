package types

import (
	"testing"
)

func TestLanguageResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "language_id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(LanguageResponse))
}
