package types

import (
	"testing"
)

func TestSkipTurnToolResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "skip_turn_success",
  "status": "success",
  "reason": "value"
}`
	assertRoundTrip(t, wire, new(SkipTurnToolResponseModel))
}

func TestSkipTurnToolResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "skip_turn_success",
  "status": "success",
  "reason": "value"
}`
	assertRoundTrip(t, wire, new(SkipTurnToolResponseModel))
}
