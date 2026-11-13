package projects

import (
	"testing"
)

func TestBodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "from_url": "value",
  "from_content_json": "value",
  "auto_convert": true
}`
	assertRoundTrip(t, wire, new(BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost))
}
