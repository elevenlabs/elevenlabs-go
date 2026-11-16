package types

import (
	"testing"
)

func TestWidgetTermsTranslationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_hash": "value",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(WidgetTermsTranslation))
}
