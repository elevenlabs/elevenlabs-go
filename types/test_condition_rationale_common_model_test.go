package types

import (
	"testing"
)

func TestTestConditionRationaleCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "messages": [
    "value"
  ],
  "summary": "value"
}`
	assertRoundTrip(t, wire, new(TestConditionRationaleCommonModel))
}
