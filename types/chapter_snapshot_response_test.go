package types

import (
	"testing"
)

func TestChapterSnapshotResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chapter_snapshot_id": "value",
  "project_id": "value",
  "chapter_id": "value",
  "created_at_unix": 123,
  "name": "value"
}`
	assertRoundTrip(t, wire, new(ChapterSnapshotResponse))
}
