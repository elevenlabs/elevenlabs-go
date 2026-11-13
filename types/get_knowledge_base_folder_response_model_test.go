package types

import (
	"testing"
)

func TestGetKnowledgeBaseFolderResponseModelRoundTrip(t *testing.T) {
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
      "id": "value",
      "name": "value"
    }
  ],
  "type": "folder",
  "children_count": 123,
  "document_count": 123,
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  },
  "external_sync_info": {
    "type": "google_drive",
    "source_entity_id": "value",
    "integration_connection_id": "value",
    "root_folder_id": "value",
    "sync_cursor": "value",
    "last_sync_at": 123
  },
  "is_frozen": true,
  "active_sync_job": {
    "type": "google_drive",
    "folder_id": "value",
    "integration_connection_id": "value",
    "triggered_by": "on_demand",
    "status": "queued",
    "sync_type": "full",
    "items_identified": 123,
    "items_processed": 123,
    "error_message": "value",
    "started_at": 123,
    "completed_at": 123,
    "updated_at": 123,
    "id": "value",
    "created_at": 123
  }
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseFolderResponseModel))
}

func TestGetKnowledgeBaseFolderResponseModelRoundTrip2(t *testing.T) {
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
      "id": "value",
      "name": "value"
    }
  ],
  "type": "folder",
  "children_count": 123,
  "document_count": 123,
  "auto_sync_info": {
    "minimum_frequency_days": 123,
    "auto_remove": true,
    "consec_failures": 123,
    "next_refresh_by": 123
  },
  "external_sync_info": {
    "type": "google_drive",
    "source_entity_id": "value",
    "integration_connection_id": "value",
    "root_folder_id": "value",
    "sync_cursor": "value",
    "last_sync_at": 123
  },
  "is_frozen": true,
  "active_sync_job": {
    "type": "google_drive",
    "folder_id": "value",
    "integration_connection_id": "value",
    "triggered_by": "on_demand",
    "status": "queued",
    "sync_type": "full",
    "items_identified": 123,
    "items_processed": 123,
    "error_message": "value",
    "started_at": 123,
    "completed_at": 123,
    "updated_at": 123,
    "id": "value",
    "created_at": 123
  }
}`
	assertRoundTrip(t, wire, new(GetKnowledgeBaseFolderResponseModel))
}
