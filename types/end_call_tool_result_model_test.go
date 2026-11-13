package types

import (
	"testing"
)

func TestEndCallToolResultModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "end_call_success",
  "status": "success",
  "reason": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(EndCallToolResultModel))
}

func TestEndCallToolResultModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "end_call_success",
  "status": "success",
  "reason": "value",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(EndCallToolResultModel))
}
