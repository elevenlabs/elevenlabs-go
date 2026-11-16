package types

import (
	"testing"
)

func TestTestToolResultModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "testing_tool_result",
  "status": "success",
  "reason": "value",
  "to_agent": "value",
  "to_node": "value"
}`
	assertRoundTrip(t, wire, new(TestToolResultModel))
}

func TestTestToolResultModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "testing_tool_result",
  "status": "success",
  "reason": "value",
  "to_agent": "value",
  "to_node": "value"
}`
	assertRoundTrip(t, wire, new(TestToolResultModel))
}
