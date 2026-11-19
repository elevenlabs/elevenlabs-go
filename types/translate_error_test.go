package types

import (
	"testing"
)

func TestTranslateErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(TranslateError))
}
