package types

import (
	"testing"
)

func TestAgentDeploymentHistoryItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "traffic_percentage_branch_id_map": {
    "key": 9.99
  },
  "deployed_at_unix_secs": 123,
  "is_active": true,
  "source": "manual",
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  }
}`
	assertRoundTrip(t, wire, new(AgentDeploymentHistoryItem))
}
