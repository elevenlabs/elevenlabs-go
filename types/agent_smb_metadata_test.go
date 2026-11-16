package types

import (
	"testing"
)

func TestAgentSmbMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_type": "customer_facing",
  "is_ephemeral": true,
  "source_fingerprint": "value"
}`
	assertRoundTrip(t, wire, new(AgentSmbMetadata))
}
