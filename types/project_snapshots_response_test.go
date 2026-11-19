package types

import (
	"testing"
)

func TestProjectSnapshotsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "snapshots": [
    {
      "project_snapshot_id": "value",
      "project_id": "value",
      "created_at_unix": 123,
      "name": "value",
      "audio_upload": {
        "key": {}
      },
      "zip_upload": {
        "key": {}
      }
    }
  ]
}`
	assertRoundTrip(t, wire, new(ProjectSnapshotsResponse))
}
