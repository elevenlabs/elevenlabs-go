package projects

import (
	"testing"
)

func TestBodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "convert_to_mpeg": true
}`
	assertRoundTrip(t, wire, new(BodyStreamStudioProjectAudioV1StudioProjectsProjectIdSnapshotsProjectSnapshotIdStreamPost))
}
