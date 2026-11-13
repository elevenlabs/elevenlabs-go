package types

import (
	"testing"
)

func TestWidgetTextContentsTranslationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": {
    "key": "value"
  },
  "text": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(WidgetTextContentsTranslation))
}
