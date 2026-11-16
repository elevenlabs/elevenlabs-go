package types

import (
	"testing"
)

func TestAgentMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "branch_id": "value",
  "workflow_node_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(AgentMetadata))
}

func TestAgentMetadataRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "branch_id": "value",
  "workflow_node_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(AgentMetadata))
}

func TestAgentMetadataRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "branch_id": "value",
  "workflow_node_id": "value",
  "version_id": "value"
}`
	assertRoundTrip(t, wire, new(AgentMetadata))
}
