package types

import (
	"testing"
)

func TestAgentVersionMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "agent_id": "value",
  "branch_id": "value",
  "version_description": "value",
  "seq_no_in_branch": 123,
  "time_committed_secs": 123,
  "parents": {
    "in_branch_parent_id": "value",
    "out_of_branch_parent_id": "value",
    "merged_into_branch_id": "value",
    "merged_from_branch_id": "value",
    "merged_from_version_id": "value",
    "rebased_from_version_id": "value"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "merged_authors": [
    {
      "is_creator": true,
      "creator_name": "value",
      "creator_email": "value",
      "role": "admin",
      "anonymous_access_level_override": "admin",
      "access_source": "creator"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AgentVersionMetadata))
}

func TestAgentVersionMetadataRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "agent_id": "value",
  "branch_id": "value",
  "version_description": "value",
  "seq_no_in_branch": 123,
  "time_committed_secs": 123,
  "parents": {
    "in_branch_parent_id": "value",
    "out_of_branch_parent_id": "value",
    "merged_into_branch_id": "value",
    "merged_from_branch_id": "value",
    "merged_from_version_id": "value",
    "rebased_from_version_id": "value"
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "merged_authors": [
    {
      "is_creator": true,
      "creator_name": "value",
      "creator_email": "value",
      "role": "admin",
      "anonymous_access_level_override": "admin",
      "access_source": "creator"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AgentVersionMetadata))
}
