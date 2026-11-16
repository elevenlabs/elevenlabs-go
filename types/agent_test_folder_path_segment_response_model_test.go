package types

import (
	"testing"
)

func TestAgentTestFolderPathSegmentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(AgentTestFolderPathSegmentResponseModel))
}

func TestAgentTestFolderPathSegmentResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(AgentTestFolderPathSegmentResponseModel))
}
