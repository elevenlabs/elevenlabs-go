package types

import (
	"testing"
)

func TestSingleTestRunRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "test_id": "value",
  "workflow_node_id": "value",
  "root_folder_id": "value",
  "root_folder_name": "value"
}`
	assertRoundTrip(t, wire, new(SingleTestRunRequestModel))
}
