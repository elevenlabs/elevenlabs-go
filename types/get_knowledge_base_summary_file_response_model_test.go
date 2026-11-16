package types

import (
	"testing"
)

func TestGetKnowledgeBaseSummaryFileResponseModelRoundTrip(t *testing.T) {
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
  "type": "file",
  "external_sync_info": {
    "type": "google_drive",
    "source_entity_id": "value",
    "integration_connection_id": "value",
    "source_parent_entity_id": "value",
    "source_mime_type": "value",
    "source_modified_time": "2024-01-01T00:00:00Z",
    "root_folder_id": "value"
  },
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  },
  "refresh_status": {
    "status": "queued",
    "enqueued_at": 123,
    "started_at": 123,
    "completed_at": 123,
    "last_synced_at": 123,
    "error_message": "value"
  },
  "is_frozen": true
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseSummaryFileResponseModel))
}

func TestGetKnowledgeBaseSummaryFileResponseModelRoundTrip2(t *testing.T) {
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
  "type": "file",
  "external_sync_info": {
    "type": "google_drive",
    "source_entity_id": "value",
    "integration_connection_id": "value",
    "source_parent_entity_id": "value",
    "source_mime_type": "value",
    "source_modified_time": "2024-01-01T00:00:00Z",
    "root_folder_id": "value"
  },
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  },
  "refresh_status": {
    "status": "queued",
    "enqueued_at": 123,
    "started_at": 123,
    "completed_at": 123,
    "last_synced_at": 123,
    "error_message": "value"
  },
  "is_frozen": true
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseSummaryFileResponseModel))
}
