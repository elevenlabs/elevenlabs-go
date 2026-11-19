package types

import (
	"testing"
)

func TestProjectSnapshotExtendedResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "project_snapshot_id": "value",
  "project_id": "value",
  "created_at_unix": 123,
  "name": "value",
  "audio_upload": {
    "key": {}
  },
  "zip_upload": {
    "key": {}
  },
  "character_alignments": [
    {
      "characters": [
        "value"
      ],
      "character_start_times_seconds": [
        9.99
      ],
      "character_end_times_seconds": [
        9.99
      ]
    }
  ],
  "audio_duration_secs": 9.99
}`
	assertRoundTrip(t, wire, new(ProjectSnapshotExtendedResponseModel))
}
