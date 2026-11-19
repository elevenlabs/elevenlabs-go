package types

import (
	"testing"
)

func TestCaptionStyleCharacterAnimationModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enter_type": "none",
  "exit_type": "none"
}`
	assertRoundTrip(t, wire, new(CaptionStyleCharacterAnimationModel))
}
