package types

import (
	"testing"
)

func TestAgentMetadataResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123
}`
	assertRoundTrip(t, wire, new(AgentMetadataResponseModel))
}

func TestAgentMetadataResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123
}`
	assertRoundTrip(t, wire, new(AgentMetadataResponseModel))
}
