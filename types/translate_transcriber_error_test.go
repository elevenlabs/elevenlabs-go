package types

import (
	"testing"
)

func TestTranslateTranscriberErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "transcriber_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(TranslateTranscriberError))
}
