package types

import (
	"testing"
)

func TestExperimentAssignmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "key": "value",
  "variant": "value",
  "source": "server_branch",
  "experiment_id": "value"
}`
	assertRoundTrip(t, wire, new(ExperimentAssignment))
}
