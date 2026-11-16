package types

import (
	"testing"
)

func TestAgentMetadataDbModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "created_from": "cli",
  "last_updated_from": "cli"
}`
	assertRoundTrip(t, wire, new(AgentMetadataDbModel))
}
