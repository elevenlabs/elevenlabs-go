package types

import (
	"testing"
)

func TestCustomSipHeaderWithDynamicVariableRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "value": "value",
  "type": "dynamic",
  "key": "value"
}`
	assertRoundTrip(t, wire, new(CustomSipHeaderWithDynamicVariable))
}
