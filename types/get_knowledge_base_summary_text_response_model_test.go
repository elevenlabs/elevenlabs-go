package types

import (
	"testing"
)

func TestGetKnowledgeBaseSummaryTextResponseModelRoundTrip(t *testing.T) {
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
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
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
  "type": "text"
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseSummaryTextResponseModel))
}

func TestGetKnowledgeBaseSummaryTextResponseModelRoundTrip2(t *testing.T) {
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
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
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
  "type": "text"
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseSummaryTextResponseModel))
}
