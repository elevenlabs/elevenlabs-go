package types

import (
	"testing"
)

func TestKbExternalSyncJobRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(KbExternalSyncJob))
}
