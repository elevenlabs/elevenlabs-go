package types

import (
	"testing"
)

func TestTranslateStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "status",
  "status": "value"
}`
	assertRoundTrip(t, wire, new(TranslateStatus))
}
