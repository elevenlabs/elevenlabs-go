package types

import (
	"testing"
)

func TestTestConditionResultCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result": "success",
  "rationale": {
    "messages": [
      "value"
    ],
    "summary": "value"
  }
}`
	assertRoundTrip(t, wire, new(TestConditionResultCommonModel))
}
