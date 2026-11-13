package types

import (
	"testing"
)

func TestChapterSnapshotsResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "snapshots": [
    {
      "chapter_snapshot_id": "value",
      "project_id": "value",
      "chapter_id": "value",
      "created_at_unix": 123,
      "name": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(ChapterSnapshotsResponse))
}
