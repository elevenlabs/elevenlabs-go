package types

import (
	"testing"
)

func TestWorkflowToolMaxIterationsExceededStepModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "step_latency_secs": 9.99,
  "type": "max_iterations_exceeded",
  "max_iterations": 123
}`
	assertRoundTrip(t, wire, new(WorkflowToolMaxIterationsExceededStepModel))
}

func TestWorkflowToolMaxIterationsExceededStepModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "step_latency_secs": 9.99,
  "type": "max_iterations_exceeded",
  "max_iterations": 123
}`
	assertRoundTrip(t, wire, new(WorkflowToolMaxIterationsExceededStepModel))
}
