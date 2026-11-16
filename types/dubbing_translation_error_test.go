package types

import (
	"testing"
)

func TestDubbingTranslationErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "translation_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(DubbingTranslationError))
}
