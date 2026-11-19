package chapters

import (
	"testing"
)

func TestAudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "convert_to_mpeg": true
}`
	assertRoundTrip(t, wire, new(AudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost))
}
