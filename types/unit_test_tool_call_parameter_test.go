package types

import (
	"testing"
)

func TestUnitTestToolCallParameterRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "eval": {
    "type": "llm",
    "description": "value"
  },
  "path": "value"
}`
	assertRoundTrip(t, wire, new(UnitTestToolCallParameter))
}

func TestUnitTestToolCallParameterRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "eval": {
    "type": "llm",
    "description": "value"
  },
  "path": "value"
}`
	assertRoundTrip(t, wire, new(UnitTestToolCallParameter))
}
