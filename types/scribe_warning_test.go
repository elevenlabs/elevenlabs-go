package types

import (
	"testing"
)

func TestScribeWarningRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "warning",
  "warning": "value"
}`
	assertRoundTrip(t, wire, new(ScribeWarning))
}
