package types

import (
	"testing"
)

func TestGetKnowledgeBaseSummaryUrlResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "metadata": {
    "created_at_unix_secs": 123,
    "last_updated_at_unix_secs": 123,
    "size_bytes": 123
  },
  "supported_usages": [
    "prompt"
  ],
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "folder_parent_id": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ],
  "dependent_agents": [
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
  "type": "url",
  "url": "value",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseSummaryUrlResponseModel))
}

func TestGetKnowledgeBaseSummaryUrlResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "metadata": {
    "created_at_unix_secs": 123,
    "last_updated_at_unix_secs": 123,
    "size_bytes": 123
  },
  "supported_usages": [
    "prompt"
  ],
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "folder_parent_id": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ],
  "dependent_agents": [
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
  "type": "url",
  "url": "value",
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  }
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseSummaryUrlResponseModel))
}
