package types

import (
	"testing"
)

func TestExternalFileSyncInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "google_drive",
  "source_entity_id": "value",
  "integration_connection_id": "value",
  "source_parent_entity_id": "value",
  "source_mime_type": "value",
  "source_modified_time": "2024-01-01T00:00:00Z",
  "root_folder_id": "value"
}`
	assertRoundTrip(t, wire, new(ExternalFileSyncInfo))
}

func TestExternalFileSyncInfoRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "google_drive",
  "source_entity_id": "value",
  "integration_connection_id": "value",
  "source_parent_entity_id": "value",
  "source_mime_type": "value",
  "source_modified_time": "2024-01-01T00:00:00Z",
  "root_folder_id": "value"
}`
	assertRoundTrip(t, wire, new(ExternalFileSyncInfo))
}
