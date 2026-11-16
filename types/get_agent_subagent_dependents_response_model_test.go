package types

import (
	"testing"
)

func TestGetAgentSubagentDependentsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(GetAgentSubagentDependentsResponseModel))
}
