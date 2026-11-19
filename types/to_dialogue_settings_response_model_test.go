package types

import (
	"testing"
)

func TestToDialogueSettingsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "stability": 0.5,
  "similarity": 0.5,
  "clone_strength": 0.5
}`
	assertRoundTrip(t, wire, new(ToDialogueSettingsResponseModel))
}
