package types

import (
	"testing"
)

func TestProjectMutedTracksResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chapter_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ProjectMutedTracksResponseModel))
}
