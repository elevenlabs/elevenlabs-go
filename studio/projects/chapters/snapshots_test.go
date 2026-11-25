package chapters

import (
	"testing"
)

func TestBodyStreamChapterAudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "convert_to_mpeg": true
}`
	assertRoundTrip(t, wire, new(BodyStreamChapterAudioV1StudioProjectsProjectIdChaptersChapterIdSnapshotsChapterSnapshotIdStreamPost))
}
