package types

import (
	"testing"
)

func TestBehaviorOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "verbosity": "auto",
  "output_format": "mp3_22050_32",
  "interaction_budget": "realtime"
}`
	assertRoundTrip(t, wire, new(BehaviorOverride))
}
