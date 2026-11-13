package types

import (
	"testing"
)

func TestExternalFolderSyncInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "google_drive",
  "source_entity_id": "value",
  "integration_connection_id": "value",
  "root_folder_id": "value",
  "sync_cursor": "value",
  "last_sync_at": 123
}`
	assertRoundTrip(t, wire, new(ExternalFolderSyncInfo))
}

func TestExternalFolderSyncInfoRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "google_drive",
  "source_entity_id": "value",
  "integration_connection_id": "value",
  "root_folder_id": "value",
  "sync_cursor": "value",
  "last_sync_at": 123
}`
	assertRoundTrip(t, wire, new(ExternalFolderSyncInfo))
}
