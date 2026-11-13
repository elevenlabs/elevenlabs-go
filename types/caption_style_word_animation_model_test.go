package types

import (
	"testing"
)

func TestCaptionStyleWordAnimationModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enter_type": "none",
  "exit_type": "none"
}`
	assertRoundTrip(t, wire, new(CaptionStyleWordAnimationModel))
}
