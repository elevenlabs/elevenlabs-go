package types

import (
	"testing"
)

func TestTranslatedStringRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "value": "value"
}`
	assertRoundTrip(t, wire, new(TranslatedString))
}
