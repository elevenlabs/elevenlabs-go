package types

import (
	"testing"
)

func TestTranslateTranslationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "translation",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(TranslateTranslation))
}
