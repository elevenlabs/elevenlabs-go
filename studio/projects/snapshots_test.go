package projects

import (
	"testing"
)

func TestStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "convert_to_mpeg": true
}`
	assertRoundTrip(t, wire, new(StreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost))
}
