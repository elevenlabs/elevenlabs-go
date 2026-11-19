package types

import (
	"testing"
)

func TestPostDialDigitsDynamicVariableRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "value": "value",
  "type": "dynamic"
}`
	assertRoundTrip(t, wire, new(PostDialDigitsDynamicVariable))
}
