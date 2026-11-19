package types

import (
	"testing"
)

func TestChapterSnapshotExtendedResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chapter_snapshot_id": "value",
  "project_id": "value",
  "chapter_id": "value",
  "created_at_unix": 123,
  "name": "value",
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
  ]
}`
	assertRoundTrip(t, wire, new(ChapterSnapshotExtendedResponseModel))
}
