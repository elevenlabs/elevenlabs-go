package types

import (
	"testing"
)

func TestProjectCreationMetaResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "creation_progress": 9.99,
  "status": "draft",
  "type": "blank"
}`
	assertRoundTrip(t, wire, new(ProjectCreationMetaResponseModel))
}

func TestProjectCreationMetaResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "creation_progress": 9.99,
  "status": "draft",
  "type": "blank"
}`
	assertRoundTrip(t, wire, new(ProjectCreationMetaResponseModel))
}
